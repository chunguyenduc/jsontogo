package app

import (
	"github.com/spf13/cobra"
)

type Config struct {
	Input      []byte
	FileInput  string
	FileOutput string
	StructName string
}

func ParseConfig(cmd *cobra.Command, args []string) (*Config, error) {
	var (
		fileInput  string
		fileOutput string
		structName string
		err        error
	)
	if fileInput, err = cmd.Flags().GetString("file_input"); err != nil {
		return nil, err
	}

	if fileOutput, err = cmd.Flags().GetString("file_output"); err != nil {
		return nil, err
	}

	if structName, err = cmd.Flags().GetString("name"); err != nil {
		return nil, err
	}

	return &Config{
		Input:      []byte(args[len(args)-1]),
		FileInput:  fileInput,
		FileOutput: fileOutput,
		StructName: structName,
	}, nil
}
