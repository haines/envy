// Package cmd provides the application's command-line interface.
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	// Version of the envy application.
	Version string
)

var command = &cobra.Command{
	Use:     "envy",
	Version: Version,
	Short:   "Share environment variables",
	Long: `envy fetches shared environment variables.

Values are stored securely in AWS Parameter Store, and can be
saved to a local .env file or directly sourced into your shell.`,
	Run: func(command *cobra.Command, args []string) {
	},
}

func init() {
	command.SetVersionTemplate("{{.Name}} {{.Version}}\n")
}

// Execute the envy command.
func Execute() {
	err := command.Execute()
	if err != nil {
		os.Exit(1)
	}
}
