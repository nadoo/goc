package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var Env = GetEnv()

// GO Environment
var (
	GOVER   = GoVer()
	GOROOT  = Env["GOROOT"]
	GOBIN   = Env["GOBIN"]
	GOPATH  = Env["GOPATH"]
	GOPROXY = Env["GOPROXY"]
	GOMOD   = Env["GOMOD"]
)

// GoVer erturns the go version.
func GoVer() string {
	cmd := exec.Command("go", "version")
	cmd.Env = os.Environ()
	buf := &bytes.Buffer{}
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}

	return strings.TrimSpace(buf.String())
}

// GetEnv is uesd to get all go envs.
func GetEnv() map[string]string {
	env := make(map[string]string)

	cmd := exec.Command("go", "env")
	cmd.Env = os.Environ()
	buf := &bytes.Buffer{}
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}

	lines := strings.Split(buf.String(), "\n")
	for _, line := range lines {
		kv := strings.Split(strings.Replace(line, "set ", "", 1), "=")
		if len(kv) > 1 {
			env[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}

	return env
}
