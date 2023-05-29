package cmd

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/vault-client-go"
)

func Test_contains(t *testing.T) {
	type args struct {
		s   []string
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.s, tt.args.str); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loopTree(t *testing.T) {
	type args struct {
		source      *vault.Client
		destination *vault.Client
		ctx         context.Context
		engine      string
		dstengine   string
		path        string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loopTree(tt.args.source, tt.args.destination, tt.args.ctx, tt.args.engine, tt.args.dstengine, tt.args.path)
		})
	}
}

func TestVaultClient(t *testing.T) {
	type args struct {
		address string
		token   string
	}
	tests := []struct {
		name string
		args args
		want *vault.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VaultClient(tt.args.address, tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VaultClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeys(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		name     string
		args     args
		wantKeys []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotKeys := Keys(tt.args.m); !reflect.DeepEqual(gotKeys, tt.wantKeys) {
				t.Errorf("Keys() = %v, want %v", gotKeys, tt.wantKeys)
			}
		})
	}
}
