package vaultutility

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/vault-client-go"
)

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

func TestLoopTree(t *testing.T) {
	type args struct {
		source *vault.Client
		ctx    context.Context
		engine string
		path   string
		f      func(context.Context, string, string, string, map[string]interface{})
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoopTree(tt.args.source, tt.args.ctx, tt.args.engine, tt.args.path, tt.args.f)
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
