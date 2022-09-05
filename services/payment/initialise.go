package payment

import (
	"context"
	"log"

	"github.com/italorfeitosa/go-sqs-lab/services/_common/aws"
	"github.com/italorfeitosa/go-sqs-lab/services/_common/messaging"
	"github.com/italorfeitosa/go-sqs-lab/services/_common/messaging/sqs"
	"github.com/italorfeitosa/go-sqs-lab/services/checkout/config"
	"github.com/italorfeitosa/go-sqs-lab/services/checkout/env"
	"github.com/italorfeitosa/go-sqs-lab/services/checkout/event"
)

func Initialise() {
	env.Initialise()
	aws.Initialise()
	config.Initialise()

	consumer := sqs.NewConsumer[event.OrderCreated](config.SQS.LoadQueueURL(env.OrderCreatedQueueName))

	consumer.Subscribe(func(ctx context.Context, message messaging.Message[event.OrderCreated]) messaging.Message[any] {
		j, _ := message.JSON()
		log.Println(string(j))
		return messaging.NullMessage()
	})

	consumer.OnConsumerError(func(err error) { log.Println(err) })

	consumer.Start(context.Background())
}
