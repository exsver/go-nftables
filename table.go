package nftables

const (
	TableFamilyIPv4   = "ip"
	TableFamilyIPv6   = "ip6"
	TableFamilyInet   = "inet"
	TableFamilyArp    = "arp"
	TableFamilyBridge = "bridge"
	TableFamilyNetDev = "netdev"
)

type Table struct {
	Name   string
	Family string
}
