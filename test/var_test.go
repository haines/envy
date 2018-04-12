package test

import (
	"testing"

	. "github.com/haines/envy/test/helpers"
	"github.com/stretchr/testify/assert"
)

func TestVars(t *testing.T) {
	result := Envy().Stdin(`{{ var "hello" }},{{ var "world" }}`).Run("--var", "hello=👋", "--var", "world=🌍")

	assert.Equal(t, 0, result.ExitStatus)
	assert.Equal(t, "👋,🌍", result.Stdout)
	assert.Empty(t, result.Stderr)
}

func TestInvalidVar(t *testing.T) {
	result := Envy().Run("--var", "wat")

	assert.Equal(t, 1, result.ExitStatus)
	assert.Empty(t, result.Stdout)
	assert.Contains(t, result.Stderr, `expected name=value`)
}

func TestMissingVars(t *testing.T) {
	result := Envy().Stdin(`{{ var "ohno" }}`).Run()

	assert.Equal(t, 1, result.ExitStatus)
	assert.Empty(t, result.Stdout)
	assert.Contains(t, result.Stderr, `no value provided for "ohno"`)
}
