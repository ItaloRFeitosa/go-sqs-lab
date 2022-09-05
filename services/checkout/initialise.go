package checkout

import (
	"context"
	"log"

	"github.com/italorfeitosa/go-sqs-lab/services/_common/aws"
	"github.com/italorfeitosa/go-sqs-lab/services/checkout/command"
	"github.com/italorfeitosa/go-sqs-lab/services/checkout/config"
	"github.com/italorfeitosa/go-sqs-lab/services/checkout/env"
	"github.com/italorfeitosa/go-sqs-lab/services/checkout/producer"
)

func Initialise() {
	env.Initialise()
	aws.Initialise()
	config.Initialise()
	producer.Initialise()

	err := command.HandleCreateOrder(context.Background(), command.CreateOrder{
		Amount: 10000,
	})

	log.Print(err)
}
