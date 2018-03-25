package test

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"testing"
)

var envyPath string

type result struct {
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

func envy(args ...string) *result {
	command := exec.Command(envyPath, args...)

	stdout := bytes.Buffer{}
	command.Stdout = &stdout

	stderr := bytes.Buffer{}
	command.Stderr = &stderr

	err := command.Run()

	result := &result{
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
