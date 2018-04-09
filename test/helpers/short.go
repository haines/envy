package helpers

import "testing"

func SkipInShortMode(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
}
