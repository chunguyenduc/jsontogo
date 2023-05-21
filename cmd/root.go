/*
Copyright © 2023 Duc Chu nguyenducchu1999@gmail.com
*/
package cmd

import (
	"os"

	"github.com/chunguyenduc/jsontogo/app"
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "jsontogo",
		Short: "jsontogo - a CLI to convert JSON to Go struct",
		Long:  `jsontogo - a CLI to convert JSON to Go struct`,
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := app.ParseConfig(cmd, args)
			if err != nil {
				return err
			}
			structImporter := app.NewStructImporter(conf, os.Open)
			structBuilder := app.NewStructBuilder(conf)
			structExporter := app.NewStructExporter(conf, os.Create)

			app := app.NewApplication(structImporter, structBuilder, structExporter)
			return app.RunApp()
		},
	}
}

func Execute(cmd *cobra.Command) error {
	cmd.Flags().StringP("file_input", "f", "", "read input from JSON file")
	cmd.Flags().StringP("file_output", "o", "", "write output to Go file")
	cmd.Flags().StringP("name", "n", "", "name of struct")
	return cmd.Execute()
}
