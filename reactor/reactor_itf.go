package jenkobs_reactor

import (
	"github.com/streadway/amqp"
)

type ReactorAction interface {
	GetAction() *ActionInfo
	LoadAction(action *ActionInfo)
	OnMessage(message *amqp.Delivery)
}
