package config

import (
	"github.com/italorfeitosa/go-sqs-lab/pkg/aws"

	"github.com/italorfeitosa/go-sqs-lab/pkg/messaging/sqs"
)

var SQS sqs.Config

func Initialise() {
	SQS = sqs.NewConfig().WithSession(aws.Session)
}
