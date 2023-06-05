package cmd

import (
	"context"
	"testing"
)

func Test_saveSecretToFile(t *testing.T) {
	type args struct {
		ctx     context.Context
		engine  string
		path    string
		secret  string
		subkeys map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saveSecretToFile(tt.args.ctx, tt.args.engine, tt.args.path, tt.args.secret, tt.args.subkeys)
		})
	}
}
