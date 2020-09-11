package jenkobs_reactor

const (
	ACTION_TYPE_CI    = "ci"
	ACTION_TYPE_SHELL = "shell"
)

type ActionInfo struct {
	Project      string
	Package      string
	Status       string
	Architecture string
	Type         string
	Params       map[string]string
}
