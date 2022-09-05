package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type config struct {
	client   *sqs.SQS
	queueURL string
}

type Config interface {
	WithSession(sess *session.Session) config
	WithQueueURL(url string) config
	LoadQueueURL(queueName string) config
}

func NewConfig() Config {
	return config{}
}

func (c config) WithSession(sess *session.Session) config {
	c.client = sqs.New(sess)
	return c
}

func (c config) WithQueueURL(url string) config {
	c.queueURL = url
	return c
}

func (c config) LoadQueueURL(queueName string) config {
	result, _ := c.client.GetQueueUrl(&sqs.GetQueueUrlInput{QueueName: aws.String(queueName)})
	return c.WithQueueURL(*result.QueueUrl)
}
