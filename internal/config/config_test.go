package config

import (
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

func TestParseConfig(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	cmd2 := &cobra.Command{}
	cmd2.Flags().String("file_input", "", "")
	cmd2.SetArgs([]string{"file_input"})

	cmd3 := &cobra.Command{}
	cmd3.Flags().String("file_input", "", "")
	cmd3.Flags().String("file_output", "", "")
	cmd3.SetArgs([]string{"file_input", "file_output"})

	cmd4 := &cobra.Command{}
	cmd4.Flags().String("file_input", "", "")
	cmd4.Flags().String("file_output", "", "")
	cmd4.Flags().String("name", "", "")
	cmd4.SetArgs([]string{"file_input", "file_output", "name"})

	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "TC1 - error len args = 0",
			args:    args{},
			wantErr: true,
		},
		{
			name: "TC2 - error file input",
			args: args{
				cmd:  &cobra.Command{},
				args: []string{"1"},
			},
			wantErr: true,
		},
		{
			name: "TC3 - error file output",
			args: args{
				cmd:  cmd2,
				args: []string{"1"},
			},
			wantErr: true,
		},
		{
			name: "TC4 - error name",
			args: args{
				cmd:  cmd3,
				args: []string{"1"},
			},
			wantErr: true,
		},
		{
			name: "TC5 - ok",
			args: args{
				cmd:  cmd4,
				args: []string{"1"},
			},
			want: &Config{
				Input: []byte("1"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseConfig(tt.args.cmd, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
