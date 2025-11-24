package services

import (
	"errors"
	"github.com/rabbitmq/amqp091-go"
	"log/slog"
)

type RabbitMQueue struct {
	connection *amqp091.Connection
	channel    *amqp091.Channel
}

func NewRabbitMQueue(uri string) (*RabbitMQueue, error) {
	connection, err := amqp091.Dial(uri)

	if err != nil {
		slog.Error("connect to RabbitMQ fail:", err)
		return nil, err
	}
	channel, err := connection.Channel()
	if err != nil {
		slog.Error("create channel fail:", err)
		return nil, err
	}
	return &RabbitMQueue{
		connection: connection,
		channel:    channel,
	}, nil
}

func (r *RabbitMQueue) Publish(exchange string, key string, body []byte) error {
	if r.channel == nil {
		slog.Error("Error connecting to RabbitMQ")
		return errors.New("rabbitmq channel not initialized")
	}

	return r.channel.Publish(
		exchange,
		key,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
