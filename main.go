package main

import (
	"os"
	"os/exec"
	"strings"
)

func main() {
	newArgv := make([]string, 0, len(os.Args))
	for _, arg := range os.Args[1:] {
		switch {
		case strings.Contains(arg, "main"):
			newArgv = append(newArgv, arg)
		case strings.HasSuffix(arg, ".c"):
			newArgv = append(newArgv, "dummy.c")
		default:
			newArgv = append(newArgv, arg)
		}
	}
	cmd := exec.Command(newArgv[0], newArgv[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
