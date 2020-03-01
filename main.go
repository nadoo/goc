// ALL possible goos and goarch
// https://github.com/golang/go/blob/master/src/go/build/syslist.go
// const goosList = "aix android darwin dragonfly freebsd hurd illumos js linux nacl netbsd openbsd plan9 solaris windows zos "
// const goarchList = "386 amd64 amd64p32 arm armbe arm64 arm64be ppc64 ppc64le mips mipsle mips64 mips64le mips64p32 mips64p32le ppc riscv riscv64 s390 s390x sparc sparc64 wasm "

// for arm:
// https://github.com/golang/go/wiki/GoArm
//

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var version = "dev"

// GO Environment
var (
	GOVER   = Env("GOVERSION")
	GOROOT  = Env("GOROOT")
	GOBIN   = Env("GOBIN")
	GOPATH  = Env("GOPATH")
	GOPROXY = Env("GOPROXY")
	GOMOD   = Env("GOMOD")
)

var envs = []struct {
	key  string
	args []string
}{
	{"GOVERSION", []string{"version"}},
	{"GOROOT", []string{"env", "GOROOT"}},
	{"GOBIN", []string{"env", "GOBIN"}},
	{"GOPATH", []string{"env", "GOPATH"}},
	{"GOPROXY", []string{"env", "GOPROXY"}},
	{"GOMOD", []string{"env", "GOMOD"}},
}

// Env is uesd to get go env.
func Env(k string) (v string) {
	var args []string
	for _, env := range envs {
		if env.key == k {
			args = env.args
			break
		}
	}

	cmd := exec.Command("go", args...)
	cmd.Env = os.Environ()
	buf := &bytes.Buffer{}
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}

	return strings.TrimSpace(buf.String())
}

type build struct {
	cmd  string
	os   string
	arch string
	arm  string
	dir  string
	args []string
}

var builds = []build{
	{"b", runtime.GOOS, runtime.GOARCH, "", "", []string{"build", "-v", "-i"}},
	{"bd", runtime.GOOS, runtime.GOARCH, "", "", []string{"build", "-v", "-i", "-tags=dev"}},
	{"bdr", runtime.GOOS, runtime.GOARCH, "", "", []string{"build", "-v", "-i", "-tags=dev", "-race"}},
	{"bw", "windows", "amd64", "", "", []string{"build", "-v", "-i"}},
	{"bwd", "windows", "amd64", "", "", []string{"build", "-v", "-i", "-tags=dev"}},
	{"bwdr", "windows", "amd64", "", "", []string{"build", "-v", "-i", "-tags=dev", "-race"}},
	{"bl", "linux", "amd64", "", "", []string{"build", "-v", "-i"}},
	{"bld", "linux", "amd64", "", "", []string{"build", "-v", "-i", "-tags=dev"}},
	{"bldr", "linux", "amd64", "", "", []string{"build", "-v", "-i", "-tags=dev", "-race"}},
	{"bm", "darwin", "amd64", "", "", []string{"build", "-v", "-i"}},
	{"bmd", "darwin", "amd64", "", "", []string{"build", "-v", "-i", "-tags=dev"}},
	{"bmdr", "darwin", "amd64", "", "", []string{"build", "-v", "-i", "-tags=dev", "-race"}},

	{"rw", "windows", "amd64", "", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rw32", "windows", "386", "", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rl", "linux", "amd64", "", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rl32", "linux", "386", "", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rla", "linux", "arm64", "", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}}, // arm64 == arm v8
	{"rla5", "linux", "arm", "5", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rla6", "linux", "arm", "6", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rla7", "linux", "arm", "7", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rlm", "linux", "mips", "", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rlmle", "linux", "mipsle", "", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rm", "darwin", "amd64", "", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rm32", "darwin", "386", "", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},

	{"r", runtime.GOOS, runtime.GOARCH, "", "", []string{"run", "."}},
	{"i", runtime.GOOS, runtime.GOARCH, "", "", []string{"install"}},
	{"c", runtime.GOOS, runtime.GOARCH, "", "", []string{"clean"}},
}

// Build is used to execute build commands.
func Build(b build, args ...string) {
	cmd := exec.Command("go", args...)
	cmd.Env = os.Environ()
	cmd.Dir = b.dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("-Environments:")
	fmt.Println("   GOVER:\t" + runtime.Version())
	fmt.Println("   GOROOT:\t" + GOROOT)
	fmt.Println("   GOBIN:\t" + GOBIN)
	if GOPATH != "" {
		fmt.Println("   GOPATH:\t" + GOPATH)
	}
	fmt.Println("   GOPROXY:\t" + GOPROXY)
	fmt.Println("   GOMOD:\t" + GOMOD)
	fmt.Println("   GOOS:\t" + runtime.GOOS)
	fmt.Println("   GOARCH:\t" + runtime.GOARCH)

	fmt.Println("-Target:")
	cmd.Env = append(cmd.Env, "GOARCH="+b.arch)
	cmd.Env = append(cmd.Env, "GOOS="+b.os)

	fmt.Println("   GOOS:\t" + b.os)
	fmt.Println("   GOARCH:\t" + b.arch)

	if b.arm != "" {
		cmd.Env = append(cmd.Env, "GOARM="+b.arm)
		fmt.Println("   GOARM:\t" + b.arm)
	}

	fmt.Println("-Working Dir:")
	fmt.Printf("   %s\n", cmd.Dir)

	fmt.Println("-Exec program:")
	fmt.Printf("   path: %s\n", cmd.Path)
	fmt.Printf("   args: %v\n\n", cmd.Args[1:])

	fmt.Println("\n-Exec outputs:\n-----")

	if err := cmd.Run(); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}
}

func main() {
	argc := len(os.Args)
	if argc < 2 {
		showHelp()
		return
	}

	cmd := os.Args[1]
	dir, _ := os.Getwd()

	for _, b := range builds {
		if b.cmd == cmd {
			b.dir = dir
			b.args = append(b.args, os.Args[2:]...)
			Build(b, b.args...)
			return
		}
	}

	showHelp()
	fmt.Fprint(os.Stderr, "unknown command: "+cmd)
}
