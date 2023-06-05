package targz

import (
	"io"
	"testing"
)

func TestTar(t *testing.T) {
	type args struct {
		src     string
		writers []io.Writer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Tar(tt.args.src, tt.args.writers...); (err != nil) != tt.wantErr {
				t.Errorf("Tar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUntar(t *testing.T) {
	type args struct {
		dst string
		r   io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Untar(tt.args.dst, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Untar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
