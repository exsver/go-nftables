package nftables

import (
	"strconv"
	"strings"
)

func (c *Config) Insert(rule *Rule, num int) error {
	ruleArgs, err := rule.GenArgs()
	if err != nil {
		return err
	}

	args := []string{"insert", "rule", c.Chain.Table.Family, c.Chain.Table.Name, c.Chain.Name, "position", strconv.Itoa(num)}
	args = append(args, ruleArgs...)

	// logger
	c.Logger.Printf("Inserting nftables rule '%s' to chain '%s'", strings.Join(ruleArgs, " "), c.Chain.Name)

	return c.do(args)
}
