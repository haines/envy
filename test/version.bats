#!/usr/bin/env bats

load "helper"

@test '`envy --version` prints version information' {
  envy --version
  [ "$status" -eq 0 ]
  grep "^envy " <<< "$output"
}
