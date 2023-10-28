// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package auth

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/sirupsen/logrus"

	cilium "github.com/cilium/proxy/go/cilium/api"
	"github.com/cilium/proxy/proxylib/proxylib"
)

//
// R2D2 Parser
//
// This is a toy protocol to teach people how to build a Cilium golang proxy parser.
//

// Current R2D2 parser supports filtering on a basic text protocol with 4 request-types:
// "READ <filename>\r\n"  - Read a file from the Droid
// "WRITE <filename>\r\n" - Write a file to the Droid
// "HALT\r\n" - Shutdown the Droid
// "RESET\r\n" - Reset the Droid to factory settings
//
// Replies include a status of either "OK\r\n", "ERROR\r\n" for "WRITE", "HALT", or "RESET".
//  Replies for "READ" are either "OK <filedata>\r\n" or "ERROR\r\n".
//
//
// Policy Examples:
// {cmd : "READ"}  - Allow all reads, no other commands.
// {cmd : "READ", file : "/public/.*" }  - Allow reads that are in the public directory
// {file : "/public/.*" } - Allow read/write on the public directory.
// {cmd : "HALT"} - Allow shutdown, but no other actions.

type authRule struct {
	cmdExact string
	username string
}

type authRequestData struct {
	cmd      string
	username string
}

func (rule *authRule) Matches(data interface{}) bool {
	// Cast 'data' to the type we give to 'Matches()'

	reqData, ok := data.(authRequestData)
	if !ok {
		log.Printf("Matches() called with type other than authRequestData")
		return false
	}
	if len(rule.cmdExact) > 0 && rule.cmdExact != reqData.cmd {
		log.Printf("AuthRule: cmd mismatch %v, %v", rule.cmdExact, reqData.cmd)
		return false
	}
	// if rule.username == "" {
	// 	log.Fatalf("Auth rule: username should be provided %s", rule.username)
	// 	return false
	// }
	log.Printf("policy match for rule: '%v' '%v'", rule.cmdExact, rule.username)
	return true
}

// ruleParser parses protobuf L7 rules to enforcement objects
// May panic
func ruleParser(rule *cilium.PortNetworkPolicyRule) []proxylib.L7NetworkPolicyRule {
	l7Rules := rule.GetL7Rules()
	if l7Rules == nil {
		return nil
	}

	allowRules := l7Rules.GetL7AllowRules()
	rules := make([]proxylib.L7NetworkPolicyRule, 0, len(allowRules))
	for _, l7Rule := range allowRules {
		var rr authRule
		for k, v := range l7Rule.Rule {
			switch k {
			case "cmd":
				rr.cmdExact = v
			case "username":
				if v != "" {

					rr.username = v
				}
			default:
				proxylib.ParseError(fmt.Sprintf("Unsupported key: %s", k), rule)
			}
		}

		if rr.cmdExact != "" &&
			rr.cmdExact != "LOGIN" &&
			rr.cmdExact != "SIGNUP" &&
			rr.cmdExact != "LOGOUT" {
			proxylib.ParseError(fmt.Sprintf("Unable to parse L7 auth rule with invalid cmd: '%s'", rr.cmdExact), rule)
		}
		if (rr.username != "") && !(rr.cmdExact == "" || rr.cmdExact == "LOGIN" || rr.cmdExact == "SIGNUP") {
			proxylib.ParseError(fmt.Sprintf("Unable to parse L7 auth rule, cmd '%s' is not compatible with 'username'", rr.cmdExact), rule)
		}
		log.Printf("Parsed rule '%s' '%s'", rr.cmdExact, rr.username)
		rules = append(rules, &rr)
	}
	return rules
}

type factory struct{}

func init() {
	logrus.Debug("init(): Registering authRuleFactory")
	proxylib.RegisterParserFactory("auth", &factory{})
	proxylib.RegisterL7RuleParser("auth", ruleParser)
}

type parser struct {
	connection *proxylib.Connection
}

func (f *factory) Create(connection *proxylib.Connection) interface{} {
	logrus.Debugf("AuthRuleFactory: Create: %v", connection)

	return &parser{connection: connection}
}

func (p *parser) OnData(reply, endStream bool, dataArray [][]byte) (proxylib.OpType, int) {

	// inefficient, but simple
	data := string(bytes.Join(dataArray, []byte{}))
	msgLen := strings.Index(data, "\r\n")
	if msgLen < 0 {
		// No delimiter, request more data
		logrus.Debugf("No delimiter found, requesting more bytes")
		return proxylib.MORE, 1
	}

	msgStr := data[:msgLen] // read single request
	msgLen += 2             // include "\r\n"
	logrus.Debugf("Request = '%s'", msgStr)

	// we don't process reply traffic for now
	fmt.Println("Reply msg ", reply)
	if reply {
		log.Printf("reply, passing %v bytes", msgLen)
		return proxylib.PASS, msgLen
	}

	fields := strings.Split(msgStr, " ")
	if len(fields) < 1 {
		return proxylib.ERROR, int(proxylib.ERROR_INVALID_FRAME_TYPE)
	}
	reqData := authRequestData{cmd: fields[0]}
	if len(fields) == 2 {
		reqData.username = fields[1]
	}

	matches := true
	access_log_entry_type := cilium.EntryType_Request

	if !p.connection.Matches(reqData) {
		matches = false
		access_log_entry_type = cilium.EntryType_Denied
	}

	p.connection.Log(access_log_entry_type,
		&cilium.LogEntry_GenericL7{
			GenericL7: &cilium.L7LogEntry{
				Proto: "auth",
				Fields: map[string]string{
					"cmd":  reqData.cmd,
					"file": reqData.username,
				},
			},
		})

	if !matches {
		p.connection.Inject(true, []byte("ERROR\r\n"))
		logrus.Debugf("Policy mismatch, dropping %d bytes", msgLen)
		return proxylib.DROP, msgLen
	}

	return proxylib.PASS, msgLen
}
