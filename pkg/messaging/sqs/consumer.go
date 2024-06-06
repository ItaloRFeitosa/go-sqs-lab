package sqs

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/italorfeitosa/go-sqs-lab/pkg/messaging"
	"github.com/italorfeitosa/go-sqs-lab/pkg/semaphore"
)

type consumer[T any] struct {
	config                config
	semaphore             semaphore.Semaphore
	handleMessage         messaging.ConsumerHandlerFunc[T]
	replyProducer         messaging.Producer
	errorProducer         messaging.Producer
	consumerErrorHandlers []func(error)
}

const defaultMaxConcurrency = 200
const defaultMaxNumberOfMessages = 10
const defaultVisibilityTimeout = 30

func NewConsumer[T any](c config) messaging.Consumer[T] {
	sem := semaphore.New(defaultMaxConcurrency)
	return &consumer[T]{
		config:    c,
		semaphore: sem,
	}
}

func (c *consumer[T]) Subscribe(handle messaging.ConsumerHandlerFunc[T]) {
	c.handleMessage = handle
}

func (c *consumer[T]) ReplyTo(producer messaging.Producer) {
	c.replyProducer = producer
}

func (c *consumer[T]) OnError(producer messaging.Producer) {
	c.errorProducer = producer
}

func (c *consumer[T]) OnConsumerError(fn func(error)) {
	c.consumerErrorHandlers = append(c.consumerErrorHandlers, fn)
}

func (c *consumer[T]) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			out, err := c.config.client.ReceiveMessageWithContext(ctx, &sqs.ReceiveMessageInput{
				QueueUrl:            aws.String(c.config.queueURL),
				MaxNumberOfMessages: aws.Int64(defaultMaxNumberOfMessages),
				VisibilityTimeout:   aws.Int64(defaultVisibilityTimeout),
			})
			if err != nil {
				c.notifyConsumerError(fmt.Errorf("error on receive message err: %s", err))
				continue
			}
			if len(out.Messages) == 0 {
				time.Sleep(1 * time.Second)
				continue
			}
			for _, message := range out.Messages {
				ctx, cancel := context.WithCancel(ctx)
				c.semaphore.Acquire()
				go func(ctx context.Context, m *sqs.Message) {
					c.handleSQSMessage(ctx, m)
					c.semaphore.Release()
					cancel()
				}(ctx, message)
			}
		}
	}
}

func (c *consumer[T]) handleSQSMessage(ctx context.Context, m *sqs.Message) {
	message := messaging.Message[T]{}
	if err := json.Unmarshal([]byte(*m.Body), &message); err != nil {
		errorMessage := messaging.NewMessage[any](m).WithError(err)
		c.errorProducer.Publish(ctx, errorMessage)
		c.notifyConsumerError(errorMessage)
		c.deleteSQSMessage(ctx, m)
		return
	}

	replyMessage := c.handleMessage(ctx, message)

	if replyMessage.ShouldRetry() {
		return
	}

	if replyMessage.ShouldReply() {
		c.replyProducer.Publish(ctx, replyMessage)
	}

	if replyMessage.Err != nil {
		c.errorProducer.Publish(ctx, replyMessage)
	}

	c.deleteSQSMessage(ctx, m)
}

func (c *consumer[T]) deleteSQSMessage(ctx context.Context, m *sqs.Message) {
	_, err := c.config.client.DeleteMessageWithContext(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(c.config.queueURL),
		ReceiptHandle: m.ReceiptHandle,
	})
	if err != nil {
		errorMessage := messaging.NewMessage[any](m).WithError(err)
		c.notifyConsumerError(errorMessage)
	}
}

func (c *consumer[T]) notifyConsumerError(err error) {
	for _, fn := range c.consumerErrorHandlers {
		fn(err)
	}
}
