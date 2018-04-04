package envy

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func parameterGetter(config *Config) (func(string) (string, error), error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{CredentialsChainVerboseErrors: aws.Bool(true)},
		Profile:           config.Profile,
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}

	client := ssm.New(sess)

	return func(name string) (string, error) {
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
