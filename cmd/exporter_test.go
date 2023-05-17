package cmd

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStructExporter(t *testing.T) {
	e := NewStructExporter(&Config{}, nil)
	assert.NotNil(t, e)
}

func Test_structExport_Export(t *testing.T) {
	type fields struct {
		Filename      string
		ReadWriter    io.ReadWriteCloser
		ReadWriteFunc func(string) (*os.File, error)
	}
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			fields: fields{
				Filename: "file",
				ReadWriteFunc: func(s string) (*os.File, error) {
					return nil, fmt.Errorf("error")
				},
			},
			wantErr: true,
		},

		{
			fields: fields{
				Filename: "file",
				ReadWriteFunc: func(s string) (*os.File, error) {
					return os.Stdout, nil
				},
			},
		},

		{
			fields: fields{
				Filename: "",
				ReadWriteFunc: func(s string) (*os.File, error) {
					return os.Stdout, nil
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &structExport{
				Filename:      tt.fields.Filename,
				ReadWriter:    tt.fields.ReadWriter,
				ReadWriteFunc: tt.fields.ReadWriteFunc,
			}
			got, err := s.Export(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("structExport.Export() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("structExport.Export() = %v, want %v", got, tt.want)
			}
		})
	}
}
