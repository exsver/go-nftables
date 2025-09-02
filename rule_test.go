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
