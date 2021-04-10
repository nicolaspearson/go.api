package rabbitmq

import "github.com/streadway/amqp"

type IRabbitMq interface {
	Bind(exchangeName, queueName string) error
	BindQueue(queueName string, exchangeName string) error
	Close()
	Consume(queueName string, prefetchCount int, onConsumed func(message []byte) bool) error
	DeclareExchange(name string) error
	DeclareQueue(name string) (amqp.Queue, error)
	Publish(exchangeName string, body []byte) error
	Purge(queueName string) error
}

type RabbitMq struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

type Opts struct {
	Username    string
	Password    string
	Host        string
	VirtualHost string
}
