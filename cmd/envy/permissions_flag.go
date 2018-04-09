package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type permissionsValue os.FileMode

var permissionsFormat = regexp.MustCompile(`\A0?[0-7]{3}\z`)

func newPermissionsValue(value os.FileMode, target *os.FileMode) *permissionsValue {
	*target = value
	return (*permissionsValue)(target)
}

func (p *permissionsValue) Set(value string) error {
	if !permissionsFormat.MatchString(value) {
		return errors.New("expected three octal digits (leading zero optional)")
	}

	octal, err := strconv.ParseUint(value, 8, 32)
	*p = permissionsValue(octal)
	return err
}

func (p *permissionsValue) Type() string {
	return "octal"
}

func (p *permissionsValue) String() string {
	return fmt.Sprintf("0%o", *p)
}
