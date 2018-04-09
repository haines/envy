// Package envy implements functions to read and parse a template file,
// interpolate with values fetched from AWS Parameter Store, and write the
// result to another file.
package envy

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var logger = log.New(os.Stderr, "", 0)

// Config holds the parameters for Run.
type Config struct {
	InputFilename  string      // if empty or "-", the template definition is read from stdin.
	OutputFilename string      // if empty or "-", the output is written to stdout.
	Permissions    os.FileMode // the permissions to set on the output file.
	SkipChmod      bool        // if true, don't set permissions on the output file.
	Profile        string      // the AWS credentials profile to use to connect to Parameter Store.
	Region         string      // the AWS region in which to connect to Parameter Store.
}

// Run parses the template definition from the input file, executes it, and writes the result to the output file.
//
// The template has access to the following functions:
//  getParameter "/path/to/value" // fetches a value from AWS Parameter store.
//  quote "value"                 // wraps a value in single quotes, escaping embedded single quotes with "'\''".
func Run(config *Config) {
	template, err := read(config)
	if err != nil {
		logger.Fatal("Failed to read template\n", err)
	}

	err = write(template, config)
	if err != nil {
		logger.Fatal("Failed to write output\n", err)
	}
}

func read(config *Config) (*template.Template, error) {
	var (
		name string
		file *os.File
		err  error
	)

	filename := config.InputFilename

	switch filename {
	case "", "-":
		name = "stdin"
		file = os.Stdin

	default:
		name = filepath.Base(filename)
		file, err = os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	getParameter, err := parameterGetter(config)
	if err != nil {
		return nil, err
	}

	return template.New(name).Funcs(template.FuncMap{
		"getParameter": getParameter,
		"quote":        quote,
	}).Parse(string(contents))
}

func write(template *template.Template, config *Config) error {
	var (
		file *os.File
		err  error
	)

	filename := config.OutputFilename
	permissions := os.FileMode(config.Permissions)

	switch filename {
	case "", "-":
		file = os.Stdout

	default:
		file, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, permissions)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	err = template.Execute(file, "")
	if err != nil {
		return err
	}

	if file != os.Stdout && !config.SkipChmod {
		err = os.Chmod(filename, permissions)
	}

	return err
}

func quote(value string) string {
	return "'" + strings.Replace(value, "'", `'\''`, -1) + "'"
}
