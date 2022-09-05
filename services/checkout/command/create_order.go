package command

import (
	"context"
	"time"

	"github.com/italorfeitosa/go-sqs-lab/services/_common/messaging"
	"github.com/italorfeitosa/go-sqs-lab/services/_common/uuid"
	"github.com/italorfeitosa/go-sqs-lab/services/checkout/event"
	"github.com/italorfeitosa/go-sqs-lab/services/checkout/producer"
)

type CreateOrder struct {
	Amount int64
}

func HandleCreateOrder(ctx context.Context, command CreateOrder) error {
	orderCreated := event.OrderCreated{
		OrderID:   uuid.New(),
		Amount:    command.Amount,
		CreatedAt: time.Now().UTC(),
	}

	return producer.OrderCreated.Publish(ctx, messaging.NewMessage[any](orderCreated))
}
