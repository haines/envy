package envy

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

var logger = log.New(os.Stderr, "", 0)

// Run parses the template definition from the first named file, executes it, and writes the output to the second named file.
// If the input filename is "-", the template definition is read from stdin.
// If the output filename is "-", the output is written to stdout.
func Run(inputFilename string, outputFilename string) {
	template, err := read(inputFilename)
	if err != nil {
		logger.Fatal("Failed to read template\n", err)
	}

	err = write(template, outputFilename)
	if err != nil {
		logger.Fatal("Failed to write output\n", err)
	}
}

func read(filename string) (*template.Template, error) {
	var (
		name string
		file *os.File
		err  error
	)

	if filename == "-" {
		name = "stdin"
		file = os.Stdin
	} else {
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

	return template.New(name).Parse(string(contents))
}

func write(template *template.Template, filename string) error {
	var (
		file *os.File
		err  error
	)

	if filename == "-" {
		file = os.Stdout
	} else {
		file, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	return template.Execute(file, "this isn't very useful yet!")
}
