#include "network.h"

#include <bcc/proto.h>
#include <linux/pkt_cls.h>
#include <stirng.h>


#define ETH_LEN 14
#define IP_TCP 6

struct Key {
	u32 src_ip;               //source ip
	u32 dst_ip;               //destination ip
	unsigned short src_port;  //source port
	unsigned short dst_port;  //destination port
};


struct Leaf {
	int timestamp;            //timestamp in ns
};

BPF_HASH(sessions, struct Key, struct Leaf, 1024);
BPF_HASH(buf_data,struct Key,char [100]);
int http_filter(struct __sk_buff *skb) {

	u8 *cursor = 0;

	struct ethernet_t *ethernet = cursor_advance(cursor, sizeof(*ethernet));
	//filter IP packets (ethernet type = 0x0800)
	if (!(ethernet->type == 0x0800)) {
		goto DROP;
	}

	struct ip_t *ip = cursor_advance(cursor, sizeof(*ip));
	//filter TCP packets (ip next protocol = 0x06)
	if (ip->nextp != IP_TCP) {
		goto DROP;
	}

	u32  tcp_header_length = 0;
	u32  ip_header_length = 0;
	u32  payload_offset = 0;
	u32  payload_length = 0;
	struct Key 	key;
	struct Leaf zero = {0};

        //calculate ip header length
        //value to multiply * 4
        //e.g. ip->hlen = 5 ; IP Header Length = 5 x 4 byte = 20 byte
        ip_header_length = ip->hlen << 2;    //SHL 2 -> *4 multiply

        //check ip header length against minimum
        if (ip_header_length < sizeof(*ip)) {
                goto DROP;
        }

        //shift cursor forward for dynamic ip header size
        void *_ = cursor_advance(cursor, (ip_header_length-sizeof(*ip)));

	struct tcp_t *tcp = cursor_advance(cursor, sizeof(*tcp));

	//retrieve ip src/dest and port src/dest of current packet
	//and save it into struct Key
	key.dst_ip = ip->dst;
	key.src_ip = ip->src;
	key.dst_port = tcp->dst_port;
	key.src_port = tcp->src_port;

	//calculate tcp header length
	//value to multiply *4
	//e.g. tcp->offset = 5 ; TCP Header Length = 5 x 4 byte = 20 byte
	tcp_header_length = tcp->offset << 2; //SHL 2 -> *4 multiply

	//calculate payload offset and length
	payload_offset = ETH_HLEN + ip_header_length + tcp_header_length;
	payload_length = ip->tlen - ip_header_length - tcp_header_length;

	//http://stackoverflow.com/questions/25047905/http-request-minimum-size-in-bytes
	//minimum length of http request is always geater than 7 bytes
	//avoid invalid access memory
	//include empty payload
	if(payload_length < 7) {
		goto DROP;
	}

	//load first 7 byte of payload into p (payload_array)
	//direct access to skb not allowed
	unsigned long p[7];
	int i = 0;
	for (i = 0; i < 7; i++) {
		p[i] = load_byte(skb, payload_offset + i);
	}

	//find a match with an HTTP message
	//HTTP
	if ((p[0] == 'H') && (p[1] == 'T') && (p[2] == 'T') && (p[3] == 'P')) {
		goto DROP;
	}
	if ((p[0] == 'H') && (p[1] == 'T') && (p[2] == 'T') && (p[3] == 'P') && (p[4] == 'S')) {
		goto DROP;
	}
	//GET
	if ((p[0] == 'G') && (p[1] == 'E') && (p[2] == 'T')) {
		goto DROP;
	}
	//POST
	if ((p[0] == 'P') && (p[1] == 'O') && (p[2] == 'S') && (p[3] == 'T')) {
		goto DROP;
	}
	//PUT
	if ((p[0] == 'P') && (p[1] == 'U') && (p[2] == 'T')) {
		goto HTTP_MATCH;
	}
	//DELETE
	if ((p[0] == 'D') && (p[1] == 'E') && (p[2] == 'L') && (p[3] == 'E') && (p[4] == 'T') && (p[5] == 'E')) {
		goto HTTP_MATCH;
	}
	//HEAD
	if ((p[0] == 'H') && (p[1] == 'E') && (p[2] == 'A') && (p[3] == 'D')) {
		goto HTTP_MATCH;
	}

	char bufferdata[100];
	strcpy(bufferdata,buf->data);

	//no HTTP match
	//check if packet belong to an HTTP session
	struct Leaf * lookup_leaf = sessions.lookup(&key);
	if(lookup_leaf) {
		//send packet to userspace
		goto DROP;
	}
	goto DROP;

	//keep the packet and send it to userspace returning -1
	HTTP_MATCH:
	//if not already present, insert into map <Key, Leaf>
	sessions.lookup_or_try_init(&key,&zero);
	buf_data.lookup_or_try_init(&key,&bufferdata)

	//send packet to userspace returning -1
	KEEP:
	return TC_ACT_OK;

	//drop the packet returning 0
	DROP:
	return TC_ACT_SHOT;

}

// int tcpconnect(void *ctx) {
//   bpf_trace_printk("[tcpconnect]\n");
//   return 0;
// }

// int socket_filter(struct __sk_buff *skb) {
//   unsigned char *cursor = 0;

//   struct ethernet_t *ethernet = cursor_advance(cursor, sizeof(*ethernet));
//   // Look for IP packets
//   if (ethernet->type != 0x0800) {
//     return 0;
//   }

//   struct ip_t *ip = cursor_advance(cursor, sizeof(*ip));

//   if (ip->nextp == 0x01) {
//     bpf_trace_printk("[socket_filter] ICMP request for %x\n", ip->dst);
//   }

//   if (ip->nextp == 0x06) {
//     bpf_trace_printk("[socket_filter] TCP packet for %x\n", ip->dst);
//     // Send TCP packets to userspace
//     return -1;
//   }

//   return 0;
// }

// int xdp(struct xdp_md *ctx) {
//   void *data = (void *)(long)ctx->data;
//   void *data_end = (void *)(long)ctx->data_end;

//   if (is_icmp_ping_request(data, data_end)) {
//     struct iphdr *iph = data + sizeof(struct ethhdr);
//     struct icmphdr *icmp = data + sizeof(struct ethhdr) + sizeof(struct iphdr);
//     bpf_trace_printk("[xdp] ICMP request for %x type %x DROPPED\n", iph->daddr,
//                      icmp->type);
//     return XDP_DROP;
//   }

//   return XDP_PASS;
// }

// int tc_drop_ping(struct __sk_buff *skb) {
//   bpf_trace_printk("[tc] ingress got packet\n");

//   void *data = (void *)(long)skb->data;
//   void *data_end = (void *)(long)skb->data_end;

//   if (is_icmp_ping_request(data, data_end)) {
//     struct iphdr *iph = data + sizeof(struct ethhdr);
//     struct icmphdr *icmp = data + sizeof(struct ethhdr) + sizeof(struct iphdr);
//     bpf_trace_printk("[tc] ICMP request for %x type %x\n", iph->daddr,
//                      icmp->type);
//     return TC_ACT_SHOT;
//   }
//   return TC_ACT_OK;
// }

// int tc_drop(struct __sk_buff *skb) {
//  bpf_trace_printk("[tc] dropping packet\n");
//  return TC_ACT_SHOT;
// }

// int tc_pingpong(struct __sk_buff *skb) {
//   bpf_trace_printk("[tc] ingress got packet");

//   void *data = (void *)(long)skb->data;
//   void *data_end = (void *)(long)skb->data_end;

//   if (!is_icmp_ping_request(data, data_end)) {
//     bpf_trace_printk("[tc] ingress not a ping request");
//     return TC_ACT_OK;
//   }

//   struct iphdr *iph = data + sizeof(struct ethhdr);
//   struct icmphdr *icmp = data + sizeof(struct ethhdr) + sizeof(struct iphdr);
//   bpf_trace_printk("[tc] ICMP request for %x type %x\n", iph->daddr,
//                    icmp->type);

//   swap_mac_addresses(skb);
//   swap_ip_addresses(skb);

//   // Change the type of the ICMP packet to 0 (ICMP Echo Reply) (was 8 for ICMP
//   // Echo request)
//   update_icmp_type(skb, 8, 0);

//   // Redirecting the modified skb on the same interface to be transmitted
//   // again
//   bpf_clone_redirect(skb, skb->ifindex, 0);

//   // We modified the packet and redirected a clone of it, so drop this one
//   return TC_ACT_SHOT;
// }