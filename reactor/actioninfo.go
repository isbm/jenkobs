package jenkobs_reactor

// ActionTypeHTTP constant for http action
const ActionTypeHTTP string = "http"

// ActionTypeShell constant for shell action
var ActionTypeShell = "shell"

// ActionInfo object
type ActionInfo struct {
	Project      string
	Package      string
	Status       string
	Architecture string
	Type         string
	Params       map[string]interface{}
}
