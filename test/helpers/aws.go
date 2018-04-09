package helpers

import (
	"fmt"
	"os"
)

type AwsConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
}

var TestAwsConfig = AwsConfig{
	AccessKeyID:     os.Getenv("ENVY_ACCESS_KEY_ID"),
	SecretAccessKey: os.Getenv("ENVY_SECRET_ACCESS_KEY"),
	Region:          os.Getenv("ENVY_REGION"),
}

const (
	credentialsFileFormat = `[%s]
aws_access_key_id = %s
aws_secret_access_key = %s
`
	configFileFormat = `[%s]
region = %s
`
)

type awsSharedFiles struct {
	Credentials string
	Config      string
}

func WriteAwsConfig(profile string, config AwsConfig) *awsSharedFiles {
	return &awsSharedFiles{
		Credentials: TempFileContaining(fmt.Sprintf(credentialsFileFormat, profile, config.AccessKeyID, config.SecretAccessKey)),
		Config:      TempFileContaining(fmt.Sprintf(configFileFormat, profile, config.Region)),
	}
}
