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
	GOVER   = GetVer()
	GOROOT  = Env["GOROOT"]
	GOBIN   = Env["GOBIN"]
	GOPATH  = Env["GOPATH"]
	GOPROXY = Env["GOPROXY"]
	GOMOD   = Env["GOMOD"]
)

var envs = []struct {
	key  string
	args []string
}{
	{"GOROOT", []string{"env", "GOROOT"}},
	{"GOBIN", []string{"env", "GOBIN"}},
	{"GOPATH", []string{"env", "GOPATH"}},
	{"GOPROXY", []string{"env", "GOPROXY"}},
	{"GOMOD", []string{"env", "GOMOD"}},
}

// GetVer is uesd to get go version.
func GetVer() string {
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
		kv := strings.Split(line, "=")
		if len(kv) > 1 {
			env[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}

	return env
}
