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

const (
	ChainTypeFilter = "filter"
	ChainTypeNat    = "nat"
	ChainTypeRoute  = "route"
)

type Chain struct {
	Name     string
	Type     string
	Hook     string
	Priority int
	Policy   string
	Table    *Table
}

func NewChain(name, hook string, table *Table) *Chain {
	return &Chain{
		Name:     name,
		Type:     ChainTypeFilter,
		Hook:     hook,
		Priority: 0,
		Policy:   ChainPolicyAccept,
		Table:    table,
	}
}

func (c *Chain) SetType(chainType string) {
	c.Type = chainType
}

func (c *Chain) SetHook(chainHook string) {
	c.Hook = chainHook
}

func (c *Chain) SetPriority(chainPriority int) {
	c.Priority = chainPriority
}

func (c *Chain) SetPolicy(chainPolicy string) {
	c.Policy = chainPolicy
}
