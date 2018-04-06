package test

import (
	"testing"

	. "github.com/haines/envy/test/helpers"
	"github.com/stretchr/testify/assert"
)

func TestQuote(t *testing.T) {
	result := Envy().Stdin(`{{ quote "foo'bar" }}`).Run()

	assert.Equal(t, 0, result.ExitStatus)
	assert.Equal(t, `'foo'\''bar'`, result.Stdout)
	assert.Empty(t, result.Stderr)
}
