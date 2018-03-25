package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	result := envy("--version")

	assert.Equal(t, 0, result.ExitStatus)
	assert.Regexp(t, "^envy ", result.Stdout)
	assert.Empty(t, result.Stderr)
}
