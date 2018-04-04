package helpers

import (
	"fmt"
	"os"
)

var AwsConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
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

func init() {
	AwsConfig.AccessKeyID = os.Getenv("ENVY_ACCESS_KEY_ID")
	AwsConfig.SecretAccessKey = os.Getenv("ENVY_SECRET_ACCESS_KEY")
	AwsConfig.Region = os.Getenv("ENVY_REGION")
}

type awsSharedFiles struct {
	Credentials string
	Config      string
}

func WriteAwsConfig(profile string) *awsSharedFiles {
	return &awsSharedFiles{
		Credentials: TempFileContaining(fmt.Sprintf(credentialsFileFormat, profile, AwsConfig.AccessKeyID, AwsConfig.SecretAccessKey)),
		Config:      TempFileContaining(fmt.Sprintf(configFileFormat, profile, AwsConfig.Region)),
	}
}
