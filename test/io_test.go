package test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultInputAndOutput(t *testing.T) {
	result := envy().Stdin("👋🌍").Run()

	assert.Equal(t, 0, result.ExitStatus)
	assert.Equal(t, "👋🌍", result.Stdout)
	assert.Empty(t, result.Stderr)
}

func TestExplicitInputFromStdin(t *testing.T) {
	result := envy().Stdin("💃🕺").Run("--input", "-")

	assert.Equal(t, 0, result.ExitStatus)
	assert.Equal(t, "💃🕺", result.Stdout)
	assert.Empty(t, result.Stderr)
}

func TestInputFromFile(t *testing.T) {
	file := tempFileContaining("🔅🔆")

	result := envy().Run("--input", file.Name())

	assert.Equal(t, 0, result.ExitStatus)
	assert.Equal(t, "🔅🔆", result.Stdout)
	assert.Empty(t, result.Stderr)
}

func TestExplicitOutputToStdout(t *testing.T) {
	result := envy().Stdin("🙂🙃").Run("--output", "-")

	assert.Equal(t, 0, result.ExitStatus)
	assert.Equal(t, "🙂🙃", result.Stdout)
	assert.Empty(t, result.Stderr)
}

func TestOutputToNewFile(t *testing.T) {
	filename := uniqueFilename()

	result := envy().Stdin("🌩🌧").Run("--output", filename)

	file := outputFile(filename)

	assert.Equal(t, 0, result.ExitStatus)
	assert.Empty(t, result.Stdout)
	assert.Empty(t, result.Stderr)

	require.True(t, file.Exists)
	assert.Equal(t, "🌩🌧", file.Contents)
	assert.Equal(t, os.FileMode(0600), file.Permissions)
}

func TestOutputToExistingFile(t *testing.T) {
	filename := uniqueFilename()
	err := ioutil.WriteFile(filename, []byte("💣💥"), 0644)
	if err != nil {
		panic(err)
	}

	result := envy().Stdin("🤜🤛").Run("--output", filename)

	file := outputFile(filename)

	assert.Equal(t, 0, result.ExitStatus)
	assert.Empty(t, result.Stdout)
	assert.Empty(t, result.Stderr)

	require.True(t, file.Exists)
	assert.Equal(t, "🤜🤛", file.Contents)
	assert.Equal(t, os.FileMode(0644), file.Permissions)
}
