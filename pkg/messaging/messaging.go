package messaging

import (
	"context"
	"encoding/json"
	"time"

	"github.com/italorfeitosa/go-sqs-lab/pkg/uuid"
)

type Producer interface {
	Publish(context.Context, Message[any]) error
}

type Consumer[T any] interface {
	Start(context.Context)
	Subscribe(ConsumerHandlerFunc[T])
	ReplyTo(Producer)
	OnError(Producer)
	OnConsumerError(func(error))
}

type ConsumerHandlerFunc[T any] func(context.Context, Message[T]) Message[any]

type Message[T any] struct {
	ID          string    `json:"id"`
	Data        T         `json:"data"`
	Err         error     `json:"error"`
	SentAt      time.Time `json:"sentAt"`
	shouldRetry bool
	shouldReply bool
	isNil       bool
}

func NewMessage[T any](data T) Message[T] {
	return Message[T]{
		ID:     uuid.New(),
		Data:   data,
		SentAt: time.Now().UTC(),
	}
}

func (m Message[T]) JSON() ([]byte, error) {
	return json.Marshal(&m)
}

func (m Message[T]) WithError(err error) Message[T] {
	m.Err = err
	return m
}

func (m Message[T]) Retry() Message[T] {
	m.shouldRetry = true
	m.shouldReply = false
	return m
}

func (m Message[T]) ShouldRetry() bool {
	return m.shouldRetry
}

func (m Message[T]) Reply() Message[T] {
	m.shouldReply = true
	m.shouldRetry = false
	return m
}

func (m Message[T]) ShouldReply() bool {
	return m.shouldReply
}

func (m Message[T]) Error() string {
	return m.Err.Error()
}

func (m Message[T]) IsNil() bool {
	return m.isNil
}

func NullMessage() Message[any] {
	return Message[any]{isNil: true}
}
