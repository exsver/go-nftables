package nftables

import "strconv"

func (c *Config) DeleteRule(handle int) error {
	args := []string{"delete", "rule", c.Chain.Table.Family, c.Chain.Table.Name, c.Chain.Name, "handle", strconv.Itoa(handle)}

	// logger
	c.Logger.Printf("Deleting rule from '%s' table '%s' family '%s'", c.Chain.Name, c.Chain.Table.Name, c.Chain.Table.Family)

	return c.do(args)
}
