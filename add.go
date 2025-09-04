package nftables

import (
	"fmt"
	"strings"
)

func (c *Config) AddRule(rule *Rule) error {
	ruleArgs, err := rule.GenArgs()
	if err != nil {
		return err
	}

	args := []string{"add", "rule", c.Chain.Table.Family, c.Chain.Table.Name, c.Chain.Name}
	args = append(args, ruleArgs...)

	// logger
	c.Logger.Printf("Appending nftables rule '%s' to chain '%s'", strings.Join(ruleArgs, " "), c.Chain.Name)

	return c.do(args)
}

func (c *Config) AddChain() error {
	args := []string{"add", "chain", c.Chain.Table.Family, c.Chain.Table.Name, c.Chain.Name,
		fmt.Sprintf("'{ type %s hook %s priority %d; policy %s; }'", c.Chain.Type, c.Chain.Hook, c.Chain.Priority, c.Chain.Policy),
	}

	// logger
	c.Logger.Printf("Creating new chain '%s' table '%s' family '%s'", c.Chain.Name, c.Chain.Table.Name, c.Chain.Table.Family)

	return c.do(args)
}

func (c *Config) AddTable() error {
	args := []string{"add", "table", c.Chain.Table.Family, c.Chain.Table.Name}

	// logger
	c.Logger.Printf("Creating new table '%s' family '%s'", c.Chain.Table.Name, c.Chain.Table.Family)

	return c.do(args)
}
