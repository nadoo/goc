// ALL possible goos and goarch
// https://github.com/golang/go/blob/master/src/go/build/syslist.go
// const goosList = "android darwin dragonfly freebsd linux nacl netbsd openbsd plan9 solaris windows zos "
// const goarchList = "386 amd64 amd64p32 arm armbe arm64 arm64be ppc64 ppc64le mips mipsle mips64 mips64le mips64p32 mips64p32le ppc s390 s390x sparc sparc64 "

// for arm:
// https://github.com/golang/go/wiki/GoArm
//

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// GO Environment
var (
	GOROOT = os.Getenv("GOROOT")
	GOPATH = os.Getenv("GOPATH")
)

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
	{"bw", "windows", "amd64", "", "", []string{"build", "-v", "-i"}},
	{"bwd", "windows", "amd64", "", "", []string{"build", "-v", "-i", "-tags=dev"}},
	{"bl", "linux", "amd64", "", "", []string{"build", "-v", "-i"}},
	{"bld", "linux", "amd64", "", "", []string{"build", "-v", "-i", "-tags=dev"}},
	{"bm", "darwin", "amd64", "", "", []string{"build", "-v", "-i"}},

	{"rw", "windows", "amd64", "", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rw32", "windows", "386", "", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rl", "linux", "amd64", "", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rl32", "linux", "386", "", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rla", "linux", "arm", "", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rla5", "linux", "arm", "5", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rla6", "linux", "arm", "6", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rla7", "linux", "arm", "7", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rla8", "linux", "arm", "8", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rlm", "linux", "mips", "", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rlmle", "linux", "mipsle", "", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rm", "darwin", "amd64", "", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},
	{"rm32", "darwin", "386", "", "", []string{"build", "-v", "-i", "-ldflags", "-s -w"}},

	{"r", runtime.GOOS, runtime.GOARCH, "", "", []string{"run"}},
	{"i", runtime.GOOS, runtime.GOARCH, "", "", []string{"install"}},
	{"c", runtime.GOOS, runtime.GOARCH, "", "", []string{"clean"}},
}

// Build :
func Build(b build, args ...string) {
	cmd := exec.Command("go", args...)
	cmd.Env = os.Environ()
	cmd.Dir = b.dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = append(cmd.Env, "GOARCH="+b.arch)
	cmd.Env = append(cmd.Env, "GOOS="+b.os)

	if b.arm != "" {
		cmd.Env = append(cmd.Env, "GOARM="+b.arm)
	}

	fmt.Println("-Environments:")
	fmt.Println("   GOROOT:" + GOROOT)
	fmt.Println("   GOPATH:" + GOPATH)
	fmt.Println("   GOARCH:" + b.arch)
	fmt.Println("   GOOS:" + b.os)

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
