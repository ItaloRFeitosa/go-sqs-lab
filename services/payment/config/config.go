package config

import (
	"github.com/italorfeitosa/go-sqs-lab/services/_common/aws"

	"github.com/italorfeitosa/go-sqs-lab/services/_common/messaging/sqs"
)

var SQS sqs.Config

func Initialise() {
	SQS = sqs.NewConfig().WithSession(aws.Session)
}
