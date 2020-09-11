package jenkobs_reactor

type ReactorAction struct {
	Project      string
	Package      string
	Status       string
	Architecture string
	Action       map[string]interface{}
}
