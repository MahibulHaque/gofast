package executor

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ExecuteCmd(name string, args []string, dir string) error {
	command := exec.Command(name, args...)
	command.Dir = dir
	var out bytes.Buffer
	var stdErr bytes.Buffer
	command.Stdout = &out
	command.Stderr = &stdErr
	if err := command.Run(); err != nil {
		return fmt.Errorf("%v\n%v", err, stdErr.String())
	}
	return nil
}
