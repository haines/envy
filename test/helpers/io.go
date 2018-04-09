package helpers

import (
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type outputFile struct {
	Exists      bool
	Contents    string
	Permissions os.FileMode
}

func OutputFile(filename string) *outputFile {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return &outputFile{Exists: false}
	} else if err != nil {
		panic(err)
	}

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return &outputFile{
		Exists:      true,
		Contents:    string(contents),
		Permissions: info.Mode().Perm(),
	}
}

func TempFileContaining(data string) string {
	file, err := ioutil.TempFile(".", "envy_test_")
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(data)
	if err != nil {
		panic(err)
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}

	return file.Name()
}

func UniqueFilename() string {
	return uuid.New().String()
}
