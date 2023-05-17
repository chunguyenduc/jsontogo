/*
Copyright © 2023 Duc Chu nguyenducchu1999@gmail.com
*/
package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jsontogo",
	Short: "jsontogo - a CLI to convert JSON to Go struct",
	Long:  `jsontogo - a CLI to convert JSON to Go struct`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			input      []byte
			fileInput  string
			fileOutput string
			structName string
			err        error
			importFunc func(string) ([]byte, error)
		)
		if fileInput, err = cmd.Flags().GetString("file_input"); err != nil {
			return err
		}

		if fileOutput, err = cmd.Flags().GetString("file_output"); err != nil {
			return err
		}

		if structName, err = cmd.Flags().GetString("name"); err != nil {
			return err
		}

		if len(fileInput) > 0 {
			input, err = openFile(fileInput)
			if err != nil {
				return err
			}
		} else {
			input = []byte(args[len(args)-1])
		}

		structBuilder := NewStructBuilder(input, fileOutput, structName, importFunc)

		return structBuilder.Run()
	},
}

func openFile(input string) ([]byte, error) {
	if ext := filepath.Ext(input); ext != ".json" {
		return nil, errors.New("error: not JSON file")
	}
	jsonFile, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	return ioutil.ReadAll(jsonFile)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("file_input", "f", "", "read input from JSON file")
	rootCmd.Flags().StringP("file_output", "o", "", "write output to Go file")
	rootCmd.Flags().StringP("name", "n", "", "name of struct")
}
