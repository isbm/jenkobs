package jenkobs_reactor

import (
	"github.com/streadway/amqp"
)

type ReactorActionItf interface {
	GetType() string
	Load(data map[string]interface{})
	OnMessage(message *amqp.Delivery)
}
