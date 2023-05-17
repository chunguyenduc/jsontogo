package cmd

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"
)

func Test_openFile(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		TODO: Add test cases.
		{
			args: args{
				input: "data.son",
			},
			wantErr: true,
		},
		{
			args: args{
				input: "data.json",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buffer bytes.Buffer
			buffer.WriteString("fake, csv, data")
			got, err := ioutil.ReadAll(&buffer)
			if (err != nil) != tt.wantErr {
				t.Errorf("openFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("openFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
