package jenkobs_reactor

import (
	"encoding/json"

	wzlib_logger "github.com/infra-whizz/wzlib/logger"
	"github.com/streadway/amqp"
)

type ReactorAction interface {
	GetAction() *ActionInfo
	LoadAction(action *ActionInfo)
	OnMessage(message *ReactorDelivery) error
}

}
