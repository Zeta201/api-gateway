#include <linux/bpf.h>
#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/tcp.h>
#include <linux/pkt_cls.h>
#include <linux/in.h>
#include <linux/string.h>
#include <bpf/bpf_helpers.h>
#include "network.h"
#include <bcc/proto.h>


SEC("http_filter")
int http_filter(struct __sk_buff *skb) {
    void *data = (void *)(long)skb->data;
    void *data_end = (void *)(long)skb->data_end;

    struct ethhdr *eth = data;
    if (data + sizeof(struct ethhdr)  > data_end) {
        return XDP_ABORTED;
    }

    if (bpf_ntohs(eth->h_proto) != ETH_P_IP)
        return XDP_PASS;
    struct iphdr *ip = data + sizeof(struct ethhdr);
    if (data + sizeof(struct ethhdr)+sizeof(struct iphdr) > data_end) {
        return XDP_ABORTED;
    }

    if (ip->protocol != IPPROTO_TCP) {
        return XDP_PASS;
    }

    struct tcphdr *tcp = (struct tcphdr *)(ip + 1);
    if (tcp + 1 > data_end) {
        return XDP_PASS;
    }

    // Calculate the offset to the TCP payload (data after the TCP header).
    unsigned int payload_offset = sizeof(struct ethhdr) + ip->ihl * 4 + tcp->doff * 4;
    void *payload = data + payload_offset;
    unsigned int payload_length = data_end - payload;

    // Check if the payload is large enough to be an HTTP packet.
    if (payload_length < 4) {
        return XDP_PASS;
    }

    // Extract the first 4 bytes of the payload for HTTP request/response detection.
    __u32 *payload_data = (__u32 *)payload;
    __u32 http_magic = bpf_ntohl(payload_data[0]);

    if (http_magic == 0x47455420 ||  // "GET "
        http_magic == 0x504f5354 ||  // "POST"
        http_magic == 0x48545450 ||  // "HTTP"
        http_magic == 0x5054482f) {  // "HTTP/1"

        // The packet contains HTTP traffic.
        // Print the HTTP payload to the kernel trace log.
        bpf_trace_printk("HTTP Payload: %.*s\n", payload_length, payload);
    }

    return XDP_PASS;
}

char _license[] SEC("license") = "GPL";
