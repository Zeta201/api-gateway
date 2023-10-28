// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package auth

import (
	"testing"

	. "github.com/cilium/checkmate"

	"github.com/cilium/proxy/proxylib/accesslog"
	"github.com/cilium/proxy/proxylib/proxylib"
	"github.com/cilium/proxy/proxylib/test"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	// logging.ToggleDebugLogs(true)
	// log.SetLevel(log.DebugLevel)

	TestingT(t)
}

type AuthSuite struct {
	logServer *test.AccessLogServer
	ins       *proxylib.Instance
}

var _ = Suite(&AuthSuite{})

// Set up access log server and Library instance for all the test cases
func (s *AuthSuite) SetUpSuite(c *C) {
	s.logServer = test.StartAccessLogServer("access_log.sock", 10)
	c.Assert(s.logServer, Not(IsNil))
	s.ins = proxylib.NewInstance("node1", accesslog.NewClient(s.logServer.Path))
	c.Assert(s.ins, Not(IsNil))
}

func (s *AuthSuite) TearDownTest(c *C) {
	s.logServer.Clear()
}

func (s *AuthSuite) TearDownSuite(c *C) {
	s.logServer.Close()
}

// func (s *AuthSuite) TestAuthOnDataIncomplete(c *C) {
// 	conn := s.ins.CheckNewConnectionOK(c, "auth", true, 1, 2, "1.1.1.1:34567", "10.0.0.2:80", "no-policy")
// 	data := [][]byte{[]byte("LOGIN john")}
// 	conn.CheckOnDataOK(c, false, false, &data, []byte{}, proxylib.MORE, 1)
// }

// func (s *AuthSuite) TestAuthOnDataBasicPass(c *C) {

// 	// allow all rule
// 	s.ins.CheckInsertPolicyText(c, "1", []string{`
// 		endpoint_ips: "1.1.1.1"
// 		endpoint_id: 2
// 		ingress_per_port_policies: <
// 		  port: 80
// 		  rules: <
// 		    l7_proto: "auth"
// 		  >
// 		>
// 		`})
// 	conn := s.ins.CheckNewConnectionOK(c, "auth", true, 1, 2, "1.1.1.1:34567", "10.0.0.2:80", "1.1.1.1")
// 	msg1 := "LOGIN sssss\r\n"
// 	msg2 := "SIGNUP sssss\r\n"
// 	msg3 := "LOGOUT\r\n"
// 	// msg4 := "RESET\r\n"
// 	data := [][]byte{[]byte(msg1 + msg2 + msg3)}
// 	conn.CheckOnDataOK(c, true, false, &data, []byte{},
// 		proxylib.PASS, len(msg1),
// 		proxylib.PASS, len(msg2),
// 		proxylib.PASS, len(msg3),
// 		// proxylib.PASS, len(msg4),
// 		proxylib.MORE, 1)
// }

// func (s *AuthSuite) TestAuthOnDataMultipleReq(c *C) {

// 	// allow all rule
// 	s.ins.CheckInsertPolicyText(c, "1", []string{`
// 		endpoint_ips: "1.1.1.1"
// 		endpoint_id: 2
// 		ingress_per_port_policies: <
// 		  port: 80
// 		  rules: <
// 		    l7_proto: "auth"
// 		  >
// 		>
// 		`})
// 	conn := s.ins.CheckNewConnectionOK(c, "auth", true, 1, 2, "1.1.1.1:34567", "10.0.0.2:80", "1.1.1.1")
// 	msg1Part1 := "LOG"
// 	msg1Part2 := "IN\r\n"
// 	data := [][]byte{[]byte(msg1Part1), []byte(msg1Part2)}
// 	conn.CheckOnDataOK(c, false, false, &data, []byte{},
// 		proxylib.PASS, len(msg1Part1+msg1Part2),
// 		proxylib.MORE, 1)
// }

func (s *AuthSuite) TestAuthOnDataAllowDenyCmd(c *C) {

	s.ins.CheckInsertPolicyText(c, "1", []string{`
		endpoint_ips: "1.1.1.1"
		endpoint_id: 2
		ingress_per_port_policies: <
		  port: 80
		  rules: <
		    l7_proto: "auth"
		    l7_rules: <
		      l7_allow_rules: <
			rule: <
			  key: "cmd"
			  value: "LOGIN"
			>
		      >
		    >
		  >
		>
		`})
	conn := s.ins.CheckNewConnectionOK(c, "auth", true, 1, 2, "1.1.1.1:34567", "10.0.0.2:80", "1.1.1.1")
	msg1 := "LOGIN xssss\r\n"
	msg2 := "SIGNUP xssss\r\n"
	data := [][]byte{[]byte(msg1 + msg2)}
	conn.CheckOnDataOK(c, true, false, &data, []byte("ERROR\r\n"),
		proxylib.PASS, len(msg1),
		proxylib.DROP, len(msg2),
		proxylib.MORE, 1)
}

// func (s *AuthSuite) TestAuthOnDataAllowDenyRegex(c *C) {

// 	s.ins.CheckInsertPolicyText(c, "1", []string{`
//     name: "cp2"
//     policy: 2
//     ingress_per_port_policies: <
//       port: 80
//       rules: <
//         l7_proto: "auth"
//         l7_rules: <
//         rule: <
//           key: "cmd"
//           value: "LOGIN"
//         >
//         rule: <
//           key: "username"
//           value: "john"
//         >
//           >
//         >
//       >
//     >
//     `})
// 	conn := s.ins.CheckNewConnectionOK(c, "auth", true, 1, 2, "1.1.1.1:34567", "10.0.0.2:80", "1.1.1.1")
// 	msg1 := "LOGIN john\r\n"
// 	// msg2 := "SIGNUP john\r\n"
// 	data := [][]byte{[]byte(msg1)}
// 	conn.CheckOnDataOK(c, false, false, &data, []byte("ERROR\r\n"),
// 		proxylib.PASS, len(msg1),
// 		// proxylib.PASS, len(msg2),
// 		proxylib.MORE, 1)
// }
