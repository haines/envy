package test

import (
	"testing"

	. "github.com/haines/envy/test/helpers"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	result := Envy().Run("--version")

	assert.Equal(t, 0, result.ExitStatus)
	assert.Regexp(t, "^envy ", result.Stdout)
	assert.Empty(t, result.Stderr)
}
