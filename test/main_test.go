package test

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"testing"

	"github.com/google/uuid"
)

var envyPath string

type envyParams struct {
	stdin string
}

type envyResult struct {
	ExitStatus int
	Stdout     string
	Stderr     string
}

func exitStatus(err error) (int, error) {
	if exiterr, ok := err.(*exec.ExitError); ok {
		if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
			return status.ExitStatus(), nil
		}
	}
	return -1, err
}

func envy() *envyParams {
	return &envyParams{}
}

func (p *envyParams) Stdin(stdin string) *envyParams {
	p.stdin = stdin
	return p
}

func (p *envyParams) Run(args ...string) *envyResult {
	command := exec.Command(envyPath, args...)

	command.Stdin = strings.NewReader(p.stdin)

	stdout := bytes.Buffer{}
	command.Stdout = &stdout

	stderr := bytes.Buffer{}
	command.Stderr = &stderr

	err := command.Run()

	result := &envyResult{
		ExitStatus: 0,
		Stdout:     stdout.String(),
		Stderr:     stderr.String(),
	}

	if err != nil {
		result.ExitStatus, err = exitStatus(err)
		if err != nil {
			panic(err)
		}
	}

	return result
}

func tempFileContaining(data string) *os.File {
	file, err := ioutil.TempFile(".", "envy_test_")
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(data)
	if err != nil {
		panic(err)
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}

	return file
}

func uniqueFilename() string {
	return uuid.New().String()
}

type fileDetails struct {
	Exists      bool
	Contents    string
	Permissions os.FileMode
}

func outputFile(filename string) *fileDetails {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return &fileDetails{Exists: false}
	} else if err != nil {
		panic(err)
	}

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return &fileDetails{
		Exists:      true,
		Contents:    string(contents),
		Permissions: info.Mode().Perm(),
	}
}

func TestMain(m *testing.M) {
	var err error
	envyPath, err = filepath.Abs("../target/envy")
	if err != nil {
		panic(err)
	}

	dir, err := ioutil.TempDir("", "envy_test_")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	err = os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
