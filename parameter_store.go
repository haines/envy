package envy

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func parameterGetter(config *Config) (func(...string) (string, error), error) {
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
			return "", err
		}

		return *result.Parameter.Value, nil
	}, nil
}
