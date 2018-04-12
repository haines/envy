package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type variablesValue map[string]string

var variableFormat = regexp.MustCompile(`\A([^=]+)=(.*)\z`)

func newVariablesValue(target *map[string]string) *variablesValue {
	*target = make(map[string]string)
	return (*variablesValue)(target)
}

func (v *variablesValue) Set(value string) error {
	nameAndValue := variableFormat.FindStringSubmatch(value)
	if nameAndValue == nil {
		return errors.New("expected name=value")
	}

	(*v)[nameAndValue[1]] = nameAndValue[2]
	return nil
}

func (v *variablesValue) Type() string {
	return "name=value"
}

func (v *variablesValue) String() string {
	pairs := make([]string, 0, len(*v))
	for name, value := range *v {
		pairs = append(pairs, fmt.Sprintf("%s=%q", name, value))
	}
	return strings.Join(pairs, ",")
}
