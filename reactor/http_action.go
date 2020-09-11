package jenkobs_reactor

import "github.com/streadway/amqp"

/*
	HTTP Action calls specified URL if criterion meets
*/

// HTTPAction object
type HTTPAction struct {
	actionInfo *ActionInfo
}

// NewHTTPAction constructor
func NewHTTPAction() *HTTPAction {
	hta := new(HTTPAction)
	return hta
}

// GetAction info on request for this action object
func (hta *HTTPAction) GetAction() *ActionInfo {
	return hta.actionInfo
}

// LoadAction info
func (hta *HTTPAction) LoadAction(action *ActionInfo) {
	hta.actionInfo = action
}

// OnMessage when delivery comes, match the criteria according to the action info
func (hta *HTTPAction) OnMessage(message *amqp.Delivery) {

}
