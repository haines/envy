package main

import (
	"os"

	"github.com/haines/envy"
	"github.com/spf13/cobra"
)

var (
	config envy.Config

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
		envy.Run(&config)
	},
}

func init() {
	command.SetVersionTemplate("{{.Name}} {{.Version}}\n")

	command.Flags().StringVarP(&config.InputFilename, "input", "i", "-", `read template from file ("-" is stdin)`)
	command.Flags().StringVarP(&config.OutputFilename, "output", "o", "-", `write output to file ("-" is stdout)`)
	command.Flags().StringVar(&config.Profile, "profile", "", "use a specific profile from your AWS credential file")
}

func main() {
	err := command.Execute()
	if err != nil {
		os.Exit(1)
	}
}
