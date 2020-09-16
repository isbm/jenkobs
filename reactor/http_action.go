package jenkobs_reactor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jinzhu/copier"
)

/*
	HTTP Action calls specified URL if criterion meets
*/

// HTTPAction object
type HTTPAction struct {
	BaseAction
}

// NewHTTPAction constructor
func NewHTTPAction() *HTTPAction {
	hta := new(HTTPAction)
	return hta
}

// Perform a HTTP/S request
func (hta *HTTPAction) request(message *ReactorDelivery) error {
	query, ok := hta.GetActionInfo().Params["query"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("Query is not defined for the HTTP action watching '%s' project", hta.GetActionInfo().Project)
	}

	url, ok := query["url"].(string)
	if !ok {
		return fmt.Errorf("URL was not specified for the HTTP action watching '%s' project", hta.GetActionInfo().Project)
	}

	method, ok := query["method"].(string)
	if !ok {
		method = "get"
	}
	//params := query["params"].(map[string]interface{})

	switch strings.ToLower(method) {
	case "post":
		req, _ := json.Marshal(map[string]string{})
		resp, err := http.Post(fmt.Sprintf("http://localhost:8080%s", url), "application/json", bytes.NewBuffer(req))
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		hta.GetLogger().Debugln(string(body))
	case "get":
	default:
		return fmt.Errorf("Method %s is not supported in HTTP action", strings.ToUpper(method))
	}

	return nil
}

// MakeActionInstance creates a self-contained instance copy
func (hta *HTTPAction) MakeActionInstance() interface{} {
	action := NewHTTPAction()
	src := *hta.actionInfo
	dst := &ActionInfo{}
	copier.Copy(&dst, &src)
	action.actionInfo = dst

	return *action
}

// OnMessage when delivery comes, match the criteria according to the action info
func (hta HTTPAction) OnMessage(message *ReactorDelivery) error {
	base := hta.MakeActionInstance().(HTTPAction)
	if !message.IsValid() {
		return fmt.Errorf("Skipping invalid message")
	}

	if base.Matches(message) {
		return base.request(message)
	}

	return nil
}
