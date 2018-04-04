package test

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/haines/envy/test/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultInputAndOutput(t *testing.T) {
	result := Envy().Stdin("👋🌍").Run()

	assert.Equal(t, 0, result.ExitStatus)
	assert.Equal(t, "👋🌍", result.Stdout)
	assert.Empty(t, result.Stderr)
}

func TestExplicitInputFromStdin(t *testing.T) {
	result := Envy().Stdin("💃🕺").Run("--input", "-")

	assert.Equal(t, 0, result.ExitStatus)
	assert.Equal(t, "💃🕺", result.Stdout)
	assert.Empty(t, result.Stderr)
}

func TestInputFromFile(t *testing.T) {
	filename := TempFileContaining("🔅🔆")

	result := Envy().Run("--input", filename)

	assert.Equal(t, 0, result.ExitStatus)
	assert.Equal(t, "🔅🔆", result.Stdout)
	assert.Empty(t, result.Stderr)
}

func TestExplicitOutputToStdout(t *testing.T) {
	result := Envy().Stdin("🙂🙃").Run("--output", "-")

	assert.Equal(t, 0, result.ExitStatus)
	assert.Equal(t, "🙂🙃", result.Stdout)
	assert.Empty(t, result.Stderr)
}

func TestOutputToNewFile(t *testing.T) {
	filename := UniqueFilename()

	result := Envy().Stdin("🌩🌧").Run("--output", filename)

	file := OutputFile(filename)

	assert.Equal(t, 0, result.ExitStatus)
	assert.Empty(t, result.Stdout)
	assert.Empty(t, result.Stderr)

	require.True(t, file.Exists)
	assert.Equal(t, "🌩🌧", file.Contents)
	assert.Equal(t, os.FileMode(0600), file.Permissions)
}

func TestOutputToExistingFile(t *testing.T) {
	filename := UniqueFilename()
	err := ioutil.WriteFile(filename, []byte("💣💥"), 0644)
	if err != nil {
		panic(err)
	}

	result := Envy().Stdin("🤜🤛").Run("--output", filename)

	file := OutputFile(filename)

	assert.Equal(t, 0, result.ExitStatus)
	assert.Empty(t, result.Stdout)
	assert.Empty(t, result.Stderr)

	require.True(t, file.Exists)
	assert.Equal(t, "🤜🤛", file.Contents)
	assert.Equal(t, os.FileMode(0644), file.Permissions)
}
