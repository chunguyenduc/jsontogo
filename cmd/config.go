package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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
		input      []byte
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

	if len(fileInput) > 0 {
		input, err = openFile(fileInput)
		if err != nil {
			return nil, err
		}
	} else {
		input = []byte(args[len(args)-1])
	}

	return &Config{
		Input:      input,
		FileInput:  fileInput,
		FileOutput: fileOutput,
		StructName: structName,
	}, nil
}

func openFile(input string) ([]byte, error) {
	if ext := filepath.Ext(input); ext != ".json" {
		return nil, fmt.Errorf("open %v: not JSON file", input)
	}
	jsonFile, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	return ioutil.ReadAll(jsonFile)
}
