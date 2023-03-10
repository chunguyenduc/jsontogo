/*
Copyright © 2023 Duc Chu nguyenducchu1999@gmail.com
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jsontogo",
	Short: "jsontogo - a CLI to convert JSON to Go struct",
	Long:  `jsontogo - a CLI to convert JSON to Go struct`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		var (
			input      []byte
			fileInput  string
			fileOutput string
			structName string
			result     string
			err        error
		)
		if fileInput, err = cmd.Flags().GetString("file_input"); err != nil {
			fmt.Println(err)
		}

		if len(fileInput) > 0 {
			fileExtension := filepath.Ext(fileInput)
			if fileExtension != ".json" {
				fmt.Printf("Error: not JSON file")
			} else {
				jsonFile, err := os.Open(fileInput)
				if err != nil {
					fmt.Println(err)
				}
				defer jsonFile.Close()
				input, _ = ioutil.ReadAll(jsonFile)
			}
		} else {
			input = []byte(args[len(args)-1])
		}

		if structName, err = cmd.Flags().GetString("name"); err != nil {
			fmt.Println(err)
		}

		result, err = jsonToGo(input, structName)
		if err != nil {
			fmt.Println(err)
		}

		if fileOutput, err = cmd.Flags().GetString("file_output"); err != nil {
			fmt.Println(err)
		}
		if len(fileOutput) > 0 {
			f, err := os.Create(fileOutput)
			if err != nil {
				fmt.Println(err)
			}
			defer f.Close()

			if _, err = f.WriteString(result); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(result)

		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("author", "a", "Duc Chu", "author name for copyright attribution")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("file_input", "f", "", "Read input from JSON file")
	rootCmd.Flags().StringP("file_output", "o", "", "Write output to Go file")
	rootCmd.Flags().StringP("name", "n", "", "Name of struct")
}
