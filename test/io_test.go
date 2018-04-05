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

func TestOutputToNewFileWithDefaultPermissions(t *testing.T) {
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

func TestOutputToNewFileWithSpecificPermissions(t *testing.T) {
	filename := UniqueFilename()

	result := Envy().Stdin("🔑🔓").Run("--output", filename, "--chmod", "644")

	file := OutputFile(filename)

	assert.Equal(t, 0, result.ExitStatus)
	assert.Empty(t, result.Stdout)
	assert.Empty(t, result.Stderr)

	require.True(t, file.Exists)
	assert.Equal(t, "🔑🔓", file.Contents)
	assert.Equal(t, os.FileMode(0644), file.Permissions)
}

func TestOutputToExistingFileWithDefaultPermissions(t *testing.T) {
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
	assert.Equal(t, os.FileMode(0600), file.Permissions)
}

func TestOutputToExistingFileWithSpecificPermissions(t *testing.T) {
	filename := UniqueFilename()
	err := ioutil.WriteFile(filename, []byte("🏃💨"), 0644)
	if err != nil {
		panic(err)
	}

	result := Envy().Stdin("🚴💨").Run("--output", filename, "--chmod", "640")

	file := OutputFile(filename)

	assert.Equal(t, 0, result.ExitStatus)
	assert.Empty(t, result.Stdout)
	assert.Empty(t, result.Stderr)

	require.True(t, file.Exists)
	assert.Equal(t, "🚴💨", file.Contents)
	assert.Equal(t, os.FileMode(0640), file.Permissions)
}

func TestOutputToExistingFileWithoutModifyingPermissions(t *testing.T) {
	filename := UniqueFilename()
	err := ioutil.WriteFile(filename, []byte("😴💤"), 0644)
	if err != nil {
		panic(err)
	}

	result := Envy().Stdin("☝️✌️").Run("--output", filename, "--no-chmod")

	file := OutputFile(filename)

	assert.Equal(t, 0, result.ExitStatus)
	assert.Empty(t, result.Stdout)
	assert.Empty(t, result.Stderr)

	require.True(t, file.Exists)
	assert.Equal(t, "☝️✌️", file.Contents)
	assert.Equal(t, os.FileMode(0644), file.Permissions)
}

func TestInvalidPermissions(t *testing.T) {
	result := Envy().Run("--chmod", "789")

	assert.Equal(t, 1, result.ExitStatus)
	assert.Empty(t, result.Stdout)
	assert.Regexp(t, "expected three octal digits", result.Stderr)
}
