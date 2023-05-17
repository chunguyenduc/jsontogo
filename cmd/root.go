/*
Copyright © 2023 Duc Chu nguyenducchu1999@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jsontogo",
	Short: "jsontogo - a CLI to convert JSON to Go struct",
	Long:  `jsontogo - a CLI to convert JSON to Go struct`,
	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := ParseConfig(cmd, args)
		if err != nil {
			return err
		}
		structBuilder := NewStructBuilder(conf.Input, conf.FileInput, conf.StructName)
		structExporter := NewStructExporter(conf.FileOutput)

		app := NewApplication(structBuilder, structExporter)
		return app.RunApp()
	},
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
