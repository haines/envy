package envy

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func paramFunc(config *Config) (func(...string) (string, error), error) {
	var region *string
	if config.Region != "" {
		region = &config.Region
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			CredentialsChainVerboseErrors: aws.Bool(true),
			Region: region,
		},
		Profile:           config.Profile,
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}

	client := ssm.New(sess)

	return func(path ...string) (string, error) {
		name := strings.Join(path, "/")
		if !strings.HasPrefix(name, "/") {
			name = "/" + name
		}

		result, err := client.GetParameter(&ssm.GetParameterInput{
			Name:           aws.String(name),
			WithDecryption: aws.Bool(true),
		})
		if err != nil {
			return "", fmt.Errorf("failed to read %q from Parameter Store:\n%v", name, err)
		}

		return *result.Parameter.Value, nil
	}, nil
}

func quote(value string) string {
	return "'" + strings.Replace(value, "'", `'\''`, -1) + "'"
}

func varFunc(config *Config) func(string) (string, error) {
	return func(name string) (string, error) {
		value, ok := config.Variables[name]
		if !ok {
			return "", fmt.Errorf("no value provided for %q", name)
		}
		return value, nil
	}
}
