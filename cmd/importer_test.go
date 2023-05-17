package cmd

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStructImporter(t *testing.T) {
	got := NewStructImporter(&Config{}, nil)
	assert.NotNil(t, got)
}

func Test_structImport_Import(t *testing.T) {
	type fields struct {
		Filename   string
		Input      []byte
		Reader     io.ReadCloser
		ReaderFunc func(string) (*os.File, error)
	}

	// writer := bufio.NewWriter()
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			fields: fields{
				Filename: "data.son",
			},
			wantErr: true,
		},
		{
			fields: fields{
				Filename: "data.json",
				ReaderFunc: func(s string) (*os.File, error) {
					return nil, fmt.Errorf("error")
				},
			},
			wantErr: true,
		},
		{
			fields: fields{
				Filename: "data.json",
				ReaderFunc: func(s string) (*os.File, error) {
					return os.Stdin, nil
				},
			},
			want: []byte{},
		},
		{
			fields: fields{
				Filename: "",
				ReaderFunc: func(s string) (*os.File, error) {
					return os.Stdout, nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &structImport{
				Filename:   tt.fields.Filename,
				Input:      tt.fields.Input,
				Reader:     tt.fields.Reader,
				ReaderFunc: tt.fields.ReaderFunc,
			}
			got, err := s.Import()
			if (err != nil) != tt.wantErr {
				t.Errorf("structImport.Import() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("structImport.Import() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
