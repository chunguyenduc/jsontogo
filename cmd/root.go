/*
Copyright © 2023 Duc Chu nguyenducchu1999@gmail.com
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jsontogo",
	Short: "jsontogo - a CLI to convert JSON to Go struct",
	Long:  `jsontogo - a CLI to convert JSON to Go struct`,
	Run: func(cmd *cobra.Command, args []string) {
		var input []byte
		fileInput, err := cmd.Flags().GetString("file_input")
		if err != nil {
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
			input = []byte(args[1])
		}
		fmt.Println(input)
		result, err := jsonToGo(input)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		glog.Errorln(err)
		os.Exit(1)
	}
}

// var Verbose bool
var FileInput string

func init() {
	rootCmd.PersistentFlags().StringP("author", "a", "Duc Chu", "author name for copyright attribution")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// rootCmd.Flags().StringP("input", "", "", "Input JSON string")
	rootCmd.Flags().StringP("file_input", "f", "", "Read input from JSON file")
	rootCmd.Flags().StringP("file_output", "o", "", "Write output to Go file")

}
