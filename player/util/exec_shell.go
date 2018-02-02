package util

import (
	"os/exec"
	"bytes"
	"fmt"
)

func ExecShell(s string) (string, error) {
	fmt.Println(s)
	cmd := exec.Command("/bin/bash", "-c", s)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	return out.String(), err
}