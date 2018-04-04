package test

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
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
