# go-nftables

Go bindings for nftables

## Install

```shell
go get -u github.com/exsver/go-nftables
```

## Examples

### Append rules

```go
package main

import (
	"log"
	"os"

	"github.com/exsver/go-nftables"
)

// Create nft configuration via cmd
//
// nft add table inet filter
// nft add chain inet filter input '{ type filter hook input priority filter; }'
func main() {
	// Create table config
	table := nftables.NewTable("filter", nftables.TableFamilyInet)

	// Create chain config
	chain := nftables.NewChain("input", nftables.ChainHookInput, table)

	// Create config
	config, err := nftables.NewConfig(chain)
	if err != nil {
		log.Fatal(err)
	}

	// Optional: Set debug logger
	config.SetLogger(log.New(os.Stdout, "Debug: ", 0))

	// Prepare rules
	rules := []nftables.Rule{
		// ip saddr 192.168.0.0/24 ip protocol icmp accept comment "Allow ICMP for 192.168.0.0/24"
		{
			SAddr:    "192.168.0.0/24",
			Protocol: "icmp",
			Jump:     nftables.RuleActionAccept,
			Comment:  "Allow ICMP for 192.168.0.0/24",
		},
		// ip protocol icmp drop comment "Deny all ICMP"
		{
			Protocol: "icmp",
			Jump:     nftables.RuleActionDrop,
			Comment:  "Deny all ICMP",
		},
	}

	// Exec nftables
	for _, rule := range rules {
		err = config.AddRule(&rule)
		if err != nil {
			log.Fatal(err)
		}
	}
}
```

### Flush chain

```go
package main

import (
	"log"
	"os"

	"github.com/exsver/go-nftables"
)

// Create nft configuration via cmd
//
// nft add table inet filter
// nft add chain inet filter input '{ type filter hook input priority filter; }'
// nft add rule inet filter input ip saddr 192.168.0.0/24 ip protocol icmp accept comment \"Allow ICMP for 192.168.0.0/24\"
// nft add rule inet filter input ip protocol icmp drop comment \"Deny all ICMP\"
func main() {
	// Create table config
	table := nftables.NewTable("filter", nftables.TableFamilyInet)

	// Create chain config
	chain := nftables.NewChain("input", nftables.ChainHookInput, table)

	// Create config
	config, err := nftables.NewConfig(chain)
	if err != nil {
		log.Fatal(err)
	}

	// Optional: Set debug logger
	config.SetLogger(log.New(os.Stdout, "Debug: ", 0))

	// Exec nftables
	err = config.FlushChain()
	if err != nil {
		log.Fatal(err)
	}
}
```
