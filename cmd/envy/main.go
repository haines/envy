package main

import (
	"os"

	"github.com/haines/envy"
	"github.com/spf13/cobra"
)

var (
	inputFilename  string
	outputFilename string

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
		envy.Run(inputFilename, outputFilename)
	},
}

func init() {
	command.SetVersionTemplate("{{.Name}} {{.Version}}\n")

	command.Flags().StringVarP(&inputFilename, "input", "i", "-", `read template from file ("-" is stdin)`)
	command.Flags().StringVarP(&outputFilename, "output", "o", "-", `write output to file ("-" is stdout)`)
}

func main() {
	err := command.Execute()
	if err != nil {
		os.Exit(1)
	}
}
