package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var Session *session.Session

func Initialise() {
	Session = session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Endpoint:    aws.String("http://localhost:9324"),
			Region:      aws.String("us-west-2"),
			Credentials: credentials.NewStaticCredentials("000000000000", "123", "123"),
		},
	}))
}
