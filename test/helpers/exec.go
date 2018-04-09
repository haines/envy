package helpers

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

var envyPath string

type Vars map[string]string

type envyParams struct {
	env   []string
	stdin string
}

type envyResult struct {
	ExitStatus int
	Stdout     string
	Stderr     string
}

func init() {
	var err error
	envyPath, err = filepath.Abs("../target/envy")
	if err != nil {
		panic(err)
	}
}

func Envy() *envyParams {
	return &envyParams{}
}

func (p *envyParams) Env(env Vars) *envyParams {
	p.env = make([]string, 0, len(env))
	for key, value := range env {
		p.env = append(p.env, key+"="+value)
	}
	return p
}

func (p *envyParams) Stdin(stdin string) *envyParams {
	p.stdin = stdin
	return p
}

func (p *envyParams) Run(args ...string) *envyResult {
	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}

	command := exec.Command(envyPath, args...)
	command.Env = p.env
	command.Stdin = strings.NewReader(p.stdin)
	command.Stdout = &stdout
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

func exitStatus(err error) (int, error) {
	if exiterr, ok := err.(*exec.ExitError); ok {
		if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
			return status.ExitStatus(), nil
		}
	}
	return -1, err
}
