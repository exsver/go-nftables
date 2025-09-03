package nftables

import (
	"reflect"
	"testing"
)

func TestRule_GenArgs(t *testing.T) {
	type fields struct {
		SAddr    string
		DAddr    string
		Protocol string
		SPort    string
		DPort    string
		NoTrack  bool
		Jump     string
		Comment  string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			name: "saddr-drop-comment",
			fields: fields{
				SAddr:   "192.168.1.10/32",
				Jump:    RuleActionDrop,
				Comment: "Drop illegal traffic",
			},
			want: []string{
				"ip", "saddr", "192.168.1.10/32",
				"drop",
				"comment", "\"Drop illegal traffic\"",
			},
			wantErr: false,
		},
		{
			name: "saddr-daddr-drop",
			fields: fields{
				SAddr: "192.168.1.10/32",
				DAddr: "192.168.1.20/32",
				Jump:  RuleActionDrop,
			},
			want: []string{
				"ip", "saddr", "192.168.1.10/32",
				"ip", "daddr", "192.168.1.20/32",
				"drop",
			},
			wantErr: false,
		},
		{
			name: "protocol-notrack-comment",
			fields: fields{
				Protocol: "udp",
				NoTrack:  true,
				Comment:  "Disable conntrack for all UDP traffic",
			},
			want: []string{
				"ip", "protocol", "udp",
				"notrack",
				"comment", "\"Disable conntrack for all UDP traffic\"",
			},
			wantErr: false,
		},
		{
			name: "sport-dport-jump-comment",
			fields: fields{
				SAddr:    "192.168.1.10/32",
				DAddr:    "192.168.1.20/32",
				Protocol: "tcp",
				SPort:    "1-1024",
				DPort:    "{80,443}",
				Jump:     "drop",
				Comment:  "Drop traffic to HTTP/HTTPS from source ports 1-1024",
			},
			want: []string{
				"ip", "saddr", "192.168.1.10/32",
				"ip", "daddr", "192.168.1.20/32",
				"tcp", "sport", "1-1024",
				"tcp", "dport", "{80,443}",
				"drop",
				"comment", "\"Drop traffic to HTTP/HTTPS from source ports 1-1024\"",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rule{
				SAddr:    tt.fields.SAddr,
				DAddr:    tt.fields.DAddr,
				Protocol: tt.fields.Protocol,
				SPort:    tt.fields.SPort,
				DPort:    tt.fields.DPort,
				NoTrack:  tt.fields.NoTrack,
				Jump:     tt.fields.Jump,
				Comment:  tt.fields.Comment,
			}

			got, err := r.GenArgs()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenArgs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
