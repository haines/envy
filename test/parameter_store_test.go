package test

import (
	"testing"

	. "github.com/haines/envy/test/helpers"
	"github.com/stretchr/testify/assert"
)

const input = `{{ getParameter "/test/string" }},{{ getParameter "/test/secure-string" }}`
const output = "foo,bar"

func TestWithCredentialsAndConfigFromEnvironment(t *testing.T) {
	SkipInShortMode(t)

	result := Envy().Env(Vars{
		"AWS_ACCESS_KEY_ID":     TestAwsConfig.AccessKeyID,
		"AWS_SECRET_ACCESS_KEY": TestAwsConfig.SecretAccessKey,
		"AWS_REGION":            TestAwsConfig.Region,
	}).Stdin(input).Run()

	assert.Equal(t, 0, result.ExitStatus)
	assert.Equal(t, output, result.Stdout)
	assert.Empty(t, result.Stderr)
}

func TestWithCredentialsFromEnvironmentAndConfigFromCommandLine(t *testing.T) {
	SkipInShortMode(t)

	result := Envy().Env(Vars{
		"AWS_ACCESS_KEY_ID":     TestAwsConfig.AccessKeyID,
		"AWS_SECRET_ACCESS_KEY": TestAwsConfig.SecretAccessKey,
		"AWS_REGION":            "overridden-by-command-line",
	}).Stdin(input).Run("--region", TestAwsConfig.Region)

	assert.Equal(t, 0, result.ExitStatus)
	assert.Equal(t, output, result.Stdout)
	assert.Empty(t, result.Stderr)
}

func TestWithCredentialsAndConfigFromDefaultProfile(t *testing.T) {
	SkipInShortMode(t)

	awsConfigFiles := WriteAwsConfig("default", TestAwsConfig)

	result := Envy().Env(Vars{
		"AWS_SHARED_CREDENTIALS_FILE": awsConfigFiles.Credentials,
		"AWS_CONFIG_FILE":             awsConfigFiles.Config,
	}).Stdin(input).Run()

	assert.Equal(t, 0, result.ExitStatus)
	assert.Equal(t, output, result.Stdout)
	assert.Empty(t, result.Stderr)
}

func TestWithCredentialsAndConfigFromProfileFromEnvironment(t *testing.T) {
	SkipInShortMode(t)

	awsConfigFiles := WriteAwsConfig("not-default", TestAwsConfig)

	result := Envy().Env(Vars{
		"AWS_SHARED_CREDENTIALS_FILE": awsConfigFiles.Credentials,
		"AWS_CONFIG_FILE":             awsConfigFiles.Config,
		"AWS_PROFILE":                 "not-default",
	}).Stdin(input).Run()

	assert.Equal(t, 0, result.ExitStatus)
	assert.Equal(t, output, result.Stdout)
	assert.Empty(t, result.Stderr)
}

func TestWithCredentialsAndConfigFromProfileFromCommandLine(t *testing.T) {
	SkipInShortMode(t)

	awsConfigFiles := WriteAwsConfig("not-default", TestAwsConfig)

	result := Envy().Env(Vars{
		"AWS_SHARED_CREDENTIALS_FILE": awsConfigFiles.Credentials,
		"AWS_CONFIG_FILE":             awsConfigFiles.Config,
		"AWS_PROFILE":                 "overridden-by-command-line",
	}).Stdin(input).Run("--profile", "not-default")

	assert.Equal(t, 0, result.ExitStatus)
	assert.Equal(t, output, result.Stdout)
	assert.Empty(t, result.Stderr)
}

func TestWithCredentialsFromProfileAndConfigFromCommandLine(t *testing.T) {
	SkipInShortMode(t)

	awsConfig := TestAwsConfig
	awsConfig.Region = "overridden-by-command-line"
	awsConfigFiles := WriteAwsConfig("default", awsConfig)

	result := Envy().Env(Vars{
		"AWS_SHARED_CREDENTIALS_FILE": awsConfigFiles.Credentials,
		"AWS_CONFIG_FILE":             awsConfigFiles.Config,
	}).Stdin(input).Run("--region", TestAwsConfig.Region)

	assert.Equal(t, 0, result.ExitStatus)
	assert.Equal(t, output, result.Stdout)
	assert.Empty(t, result.Stderr)
}
