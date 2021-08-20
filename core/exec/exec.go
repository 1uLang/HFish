package exec

import (
	"os/exec"
	"bytes"
)

func Execute(shell string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", shell)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
