package nftables

func (c *Config) FlushChain() error {
	args := []string{"flush", "chain", c.Chain.Table.Family, c.Chain.Table.Name, c.Chain.Name}

	// logger
	c.Logger.Printf("Flushing chain '%s' table '%s' family '%s'", c.Chain.Name, c.Chain.Table.Name, c.Chain.Table.Family)

	return c.do(args)
}
