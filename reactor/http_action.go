package jenkobs_reactor

import "fmt"

/*
	HTTP Action calls specified URL if criterion meets
*/

// HTTPAction object
type HTTPAction struct {
	actionInfo *ActionInfo
	BaseAction
}

// NewHTTPAction constructor
func NewHTTPAction() *HTTPAction {
	hta := new(HTTPAction)
	return hta
}

// OnMessage when delivery comes, match the criteria according to the action info
func (hta *HTTPAction) OnMessage(message *ReactorDelivery) error {
	if !message.IsValid() {
		return fmt.Errorf("Skipping invalid message")
	}

	if hta.Matches(message) {
		fmt.Println("Matches")
	}

	return nil
}
