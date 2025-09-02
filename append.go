package nftables

import "strings"

func (c *Config) Append(rule *Rule) error {
	ruleArgs, err := rule.GenArgs()
	if err != nil {
		return err
	}

	args := []string{"add", "rule", c.Chain.Type, c.Chain.Name}
	args = append(args, ruleArgs...)

	// logger
	c.Logger.Printf("Appending nftables rule '%s' to chain '%s'", strings.Join(ruleArgs, " "), c.Chain.Name)

	return c.do(args)
}
