package nftables

const (
	ChainHookPrerouting  = "prerouting"  // ip/ip6/inet
	ChainHookInput       = "input"       // ip/ip6/inet/arp
	ChainHookForward     = "forward"     // ip/ip6/inet
	ChainHookOutput      = "output"      // ip/ip6/inet/arp
	ChainHookPostrouting = "postrouting" // ip/ip6/inet
	ChainHookIngress     = "ingress"     // ip/ip6/inet/netdev
	ChainHookEgress      = "egress"      // netdev
)

const (
	ChainPolicyAccept   = "accept"
	ChainPolicyDrop     = "drop"
	ChainPolicyReturn   = "return"
	ChainPolicyContinue = "continue"
)

type Chain struct {
	Name     string
	Type     string
	Hook     string
	Priority int
	Policy   string
	Table    *Table
}
