package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		v := os.Args[i]

		if strings.Contains(v, "*") {
			expArgs, err := filepath.Glob(v)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: cannot expand glob %q", os.Args[0], v)
				os.Exit(1)
			}

			args := make([]string, len(os.Args)-1+len(expArgs))
			copy(args, os.Args[:i])
			copy(args[i:], expArgs)
			copy(args[i+len(expArgs):], os.Args[i+1:])

			os.Args = args

			i += len(expArgs)
		}
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	b1 := &bytes.Buffer{}
	b2 := &bytes.Buffer{}
	cmd.Stdout = b1
	cmd.Stderr = b2

	if err := cmd.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: cannot start command %q: %s", os.Args[0], os.Args[1], err)
		os.Exit(1)
	}

	if err := cmd.Wait(); err != nil {
		if exerr, ok := err.(*exec.ExitError); ok {
			if ws, ok := exerr.Sys().(syscall.WaitStatus); ok {
				fmt.Fprint(os.Stderr, b2.String())
				os.Exit(ws.ExitStatus())
			}
		}

		fmt.Fprintf(os.Stderr, "%s: cannot run command %q: %s", os.Args[0], os.Args[1], err)
		os.Exit(1)
	}

	fmt.Print(b1.String())
}
