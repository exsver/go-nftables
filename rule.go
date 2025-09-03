package nftables

import (
	"fmt"
	"strings"
)

const (
	RuleActionAccept = "accept"
	RuleActionDrop   = "drop"
)

type Rule struct {
	// SAddr - nft ip saddr option
	SAddr string
	// DAddr - nft ip daddr option
	DAddr string
	// Protocol - nft ip protocol option
	//  - tcp
	//  - udp
	//  - icmp
	//  ...
	Protocol string
	// SPort - nft sport option
	// Valid SrcPorts values:
	//  - "80"
	//  - "{80,443}"
	//  - "1-1024"
	SPort string
	// DPort - nft dport option
	DPort string
	// NoTrack - disable conntrack
	NoTrack bool
	// Jump - action
	//  - accept
	//  - drop
	Jump    string
	Comment string
}

func (r *Rule) GenArgs() ([]string, error) {
	var args []string

	if r.SAddr != "" {
		args = append(args, "ip", "saddr", r.SAddr)
	}

	if r.DAddr != "" {
		args = append(args, "ip", "daddr", r.DAddr)
	}

	if r.Protocol != "" {
		if r.SPort == "" && r.DPort == "" {
			args = append(args, "ip", "protocol", r.Protocol)
		}
	}

	if r.SPort != "" {
		switch r.Protocol {
		case "tcp", "udp":
			args = append(args, r.Protocol, "sport", strings.TrimSpace(r.SPort))
		default:
			return nil, fmt.Errorf("protocol must be tcp or udp")
		}
	}

	if r.DPort != "" {
		switch r.Protocol {
		case "tcp", "udp":
			args = append(args, r.Protocol, "dport", strings.TrimSpace(r.DPort))
		default:
			return nil, fmt.Errorf("protocol must be tcp or udp")
		}
	}

	if r.NoTrack {
		args = append(args, "notrack")
	}

	if r.Jump != "" {
		args = append(args, r.Jump)
	}

	if !r.NoTrack && r.Jump == "" {
		return nil, fmt.Errorf("jump or notrack must be set")
	}

	if r.Comment != "" {
		args = append(args, "comment", fmt.Sprintf("\"%s\"", strings.ReplaceAll(r.Comment, "\"", "'")))
	}

	return args, nil
}
