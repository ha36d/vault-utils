package cmd

import (
	"context"
	"testing"

	"github.com/hashicorp/vault-client-go"
)

func Test_saveSecretToKv(t *testing.T) {
	type args struct {
		destination *vault.Client
		ctx         context.Context
		engine      string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saveSecretToKv(tt.args.destination, tt.args.ctx, tt.args.engine)
		})
	}
}
