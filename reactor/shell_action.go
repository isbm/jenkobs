package jenkobs_reactor

import (
	"fmt"

	"github.com/streadway/amqp"
)

/*
	Shell action calls specific commands, if criterion meets
*/

// ShellAction object
type ShellAction struct {
	actionInfo *ActionInfo
}

// NewShellAction constructor
func NewShellAction() *ShellAction {
	shellAction := new(ShellAction)
	return shellAction
}

// GetAction info on request for this action object
func (shact *ShellAction) GetAction() *ActionInfo {
	return shact.actionInfo
}

// LoadAction info
func (shact *ShellAction) LoadAction(action *ActionInfo) {
	shact.actionInfo = action
}

// OnMessage when delivery comes, match the criteria according to the action info
func (shact *ShellAction) OnMessage(message *amqp.Delivery) {
	if message.RoutingKey == shact.actionInfo.Status {
		fmt.Println(">>> ", message.RoutingKey)
		fmt.Println(string(message.Body))
		fmt.Println("---")
	}
}
