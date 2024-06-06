package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"

	"github.com/italorfeitosa/go-sqs-lab/pkg/messaging"
)

type producer struct {
	config config
}

func NewProducer(c config) messaging.Producer {
	return &producer{
		c,
	}
}

func (s *producer) Publish(ctx context.Context, data messaging.Message[any]) error {
	jsonBytes, err := data.JSON()
	if err != nil {
		return err
	}
	_, err = s.config.client.SendMessageWithContext(ctx, &sqs.SendMessageInput{
		QueueUrl:    &s.config.queueURL,
		MessageBody: aws.String(string(jsonBytes)),
	})

	return err
}
