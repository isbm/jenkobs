package jenkobs_reactor

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	wzlib_logger "github.com/infra-whizz/wzlib/logger"
	wzlib_subprocess "github.com/infra-whizz/wzlib/subprocess"
	"github.com/jinzhu/copier"
)

/*
	Shell action calls specific commands, if criterion meets.
	Command gets back the following keywords:

		project
		package
		repo
		arch
*/

// ShellAction object
type ShellAction struct {
	BaseAction
	wzlib_logger.WzLogger
}

// NewShellAction constructor
func NewShellAction() *ShellAction {
	shellAction := new(ShellAction)
	return shellAction
}

// MakeActionInstance creates a self-contained instance copy
func (shact *ShellAction) MakeActionInstance() interface{} {
	action := NewShellAction()
	src := *shact.actionInfo
	dst := &ActionInfo{}
	if err := copier.Copy(&dst, &src); err != nil {
		shact.GetLogger().Errorf("Error copying contents of a shell instance: %s", err.Error())
	}
	action.actionInfo = dst

	return *action
}

// Format command from the parameters.
func (shact *ShellAction) formatCommand(message *ReactorDelivery) []string {
	cmdTpl, ex := shact.actionInfo.Params["command"]
	out := make([]string, 0)
	if ex {
		fmt := strings.NewReplacer(
			"{project}", message.GetProjectName(),
			"{package}", message.GetPackageName(),
			"{arch}", message.GetArch(),
			"{repo}", message.GetRepoName(),
		)
		for _, p := range cmdTpl.([]string) {
			out = append(out, fmt.Replace(p))
		}
	}
	return out
}

func (shact *ShellAction) callShellCommand(name string, args ...string) error {
	shact.GetLogger().Debugf("Calling shell command '%v' with the following parameters: %s", name, spew.Sdump(args))
	buff, err := wzlib_subprocess.BufferedExec(name, args...)
	if err != nil {
		return err
	}
	sout, serr := buff.StdoutString(), buff.StderrString()
	if err := buff.Wait(); err != nil {
		return err
	}

	if serr != "" {
		shact.GetLogger().Errorf("Error calling command %s: %s", name, serr)
	}
	if sout != "" {
		shact.GetLogger().Debugf("Command output: %s", sout)
	}

	return nil
}

// OnMessage when delivery comes, match the criteria according to the action info
func (shact *ShellAction) OnMessage(message *ReactorDelivery) error {
	if !message.IsValid() {
		return fmt.Errorf("Skipping invalid message")
	}

	action := shact.MakeActionInstance().(ShellAction)

	if action.Matches(message) {
		cmd := action.formatCommand(message)
		if len(cmd) > 0 {
			return action.callShellCommand(cmd[0], cmd[1:]...)
		}
	}

	return nil
}
