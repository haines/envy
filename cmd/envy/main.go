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

	flags := command.Flags()
	flags.Var(newVariablesValue(&config.Variables), "var", `a variable to be interpolated into the template with {{ var "name" }}`)
	flags.StringVarP(&config.InputFilename, "input", "i", "-", `read template from file ("-" is stdin)`)
	flags.StringVarP(&config.OutputFilename, "output", "o", "-", `write output to file ("-" is stdout)`)
	flags.Var(newPermissionsValue(0600, &config.Permissions), "chmod", "output file permissions")
	flags.BoolVar(&config.SkipChmod, "no-chmod", false, "don't modify output file permissions")
	flags.StringVar(&config.Profile, "profile", "", "use a specific profile from your AWS credential file")
	flags.StringVar(&config.Region, "region", "", "the AWS region to connect to")
	flags.SortFlags = false
}

func main() {
	err := command.Execute()
	if err != nil {
		os.Exit(1)
	}
}
