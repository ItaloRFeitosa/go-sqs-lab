package producer

import (
	"github.com/italorfeitosa/go-sqs-lab/pkg/messaging"
	"github.com/italorfeitosa/go-sqs-lab/pkg/messaging/sqs"
	"github.com/italorfeitosa/go-sqs-lab/services/checkout/config"
	"github.com/italorfeitosa/go-sqs-lab/services/checkout/env"
)

var OrderCreated messaging.Producer

func Initialise() {
	OrderCreated = sqs.NewProducer(config.SQS.LoadQueueURL(env.OrderCreatedQueueName))
}
