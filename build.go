package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type build struct {
	cmd  string
	os   string
	arch string
	ver  string // version of arm or amd64
	dir  string
	args []string
	help string
}

// builds uses the current version when constructing release arguments.
func builds() []build {
	return []build{
		{"b", runtime.GOOS, runtime.GOARCH, "", "", []string{"build", "-v", "-ldflags", "-X main.version=" + version}, "build package"},
		{"bd", runtime.GOOS, runtime.GOARCH, "", "", []string{"build", "-v", "-tags=dev", "-ldflags", "-X main.version=" + version}, "build dev package(-tags=dev)"},
		{"bdr", runtime.GOOS, runtime.GOARCH, "", "", []string{"build", "-v", "-tags=dev", "-race", "-ldflags", "-X main.version=" + version}, "build dev package(-tags=dev and -race)"},
		{"bw", "windows", "amd64", "", "", []string{"build", "-v", "-ldflags", "-X main.version=" + version}, "build windows package(amd64)"},
		{"bwd", "windows", "amd64", "", "", []string{"build", "-v", "-tags=dev", "-ldflags", "-X main.version=" + version}, "build windows dev package(-tags=dev)(amd64)"},
		{"bwdr", "windows", "amd64", "", "", []string{"build", "-v", "-tags=dev", "-race", "-ldflags", "-X main.version=" + version}, "build windows dev package(-tags=dev and -race)(amd64)"},
		{"bl", "linux", "amd64", "", "", []string{"build", "-v", "-ldflags", "-X main.version=" + version}, "build linux package(amd64)"},
		{"bl3", "linux", "amd64", "v3", "", []string{"build", "-v", "-ldflags", "-X main.version=" + version}, "build linux package(amd64v3)"},
		{"bld", "linux", "amd64", "", "", []string{"build", "-v", "-tags=dev", "-ldflags", "-X main.version=" + version}, "build linux dev package(-tags=dev)(amd64)"},
		{"bld3", "linux", "amd64", "v3", "", []string{"build", "-v", "-tags=dev", "-ldflags", "-X main.version=" + version}, "build linux dev package(-tags=dev)(amd64v3)"},
		{"bldr", "linux", "amd64", "", "", []string{"build", "-v", "-tags=dev", "-race", "-ldflags", "-X main.version=" + version}, "build linux dev package(-tags=dev and -race)(amd64)"},
		{"blad", "linux", "arm64", "", "", []string{"build", "-v", "-tags=dev", "-ldflags", "-X main.version=" + version}, "build linux dev package(-tags=dev)(arm64)"}, // arm64 == arm v8
		{"bladr", "linux", "arm64", "", "", []string{"build", "-v", "-tags=dev", "-race", "-ldflags", "-X main.version=" + version}, "build linux dev package(-tags=dev and -race)(arm64)"},
		{"bm", "darwin", "arm64", "", "", []string{"build", "-v", "-ldflags", "-X main.version=" + version}, "build mac package(arm64)"},
		{"bmd", "darwin", "arm64", "", "", []string{"build", "-v", "-tags=dev", "-ldflags", "-X main.version=" + version}, "build mac dev package(-tags=dev)(arm64)"},
		{"bmdr", "darwin", "arm64", "", "", []string{"build", "-v", "-tags=dev", "-race", "-ldflags", "-X main.version=" + version}, "build mac dev package(-tags=dev and -race)(arm64)"},

		{"rw", "windows", "amd64", "", "", []string{"build", "-v", "-trimpath", "-ldflags", "-s -w -X main.version=" + version}, "release windows package(amd64)"},
		{"rw3", "windows", "amd64", "v3", "", []string{"build", "-v", "-trimpath", "-ldflags", "-s -w -X main.version=" + version}, "release windows package(amd64v3)"},
		{"rw32", "windows", "386", "", "", []string{"build", "-v", "-trimpath", "-ldflags", "-s -w -X main.version=" + version}, "release windows package(x86)"},
		{"rl", "linux", "amd64", "", "", []string{"build", "-v", "-trimpath", "-ldflags", "-s -w -X main.version=" + version}, "release linux package(amd64)"},
		{"rl3", "linux", "amd64", "v3", "", []string{"build", "-v", "-trimpath", "-ldflags", "-s -w -X main.version=" + version}, "release linux package(amd64v3)"},
		{"rl32", "linux", "386", "", "", []string{"build", "-v", "-trimpath", "-ldflags", "-s -w -X main.version=" + version}, "release linux package(x86)"},
		{"rla", "linux", "arm64", "", "", []string{"build", "-v", "-trimpath", "-ldflags", "-s -w -X main.version=" + version}, "release linux package(arm64)"}, // arm64 == arm v8
		{"rla5", "linux", "arm", "5", "", []string{"build", "-v", "-trimpath", "-ldflags", "-s -w -X main.version=" + version}, "release linux package(arm v5)"},
		{"rla6", "linux", "arm", "6", "", []string{"build", "-v", "-trimpath", "-ldflags", "-s -w -X main.version=" + version}, "release linux package(arm v6)"},
		{"rla7", "linux", "arm", "7", "", []string{"build", "-v", "-trimpath", "-ldflags", "-s -w -X main.version=" + version}, "release linux package(arm v7)"},
		{"rlm", "linux", "mips", "", "", []string{"build", "-v", "-trimpath", "-ldflags", "-s -w -X main.version=" + version}, "release linux package(mips)"},
		{"rlmle", "linux", "mipsle", "", "", []string{"build", "-v", "-trimpath", "-ldflags", "-s -w -X main.version=" + version}, "release linux package(mipsle)"},
		{"rm", "darwin", "arm64", "", "", []string{"build", "-v", "-trimpath", "-ldflags", "-s -w -X main.version=" + version}, "release mac package(arm64)"},

		{"r", runtime.GOOS, runtime.GOARCH, "", "", []string{"run", "."}, "run current package"},
		{"i", runtime.GOOS, runtime.GOARCH, "", "", []string{"install", "-ldflags", "-X main.version=" + version}, "install package to `GOBIN` or `GOPATH/bin`"},
		{"c", runtime.GOOS, runtime.GOARCH, "", "", []string{"clean"}, "clean package"},
	}
}

// Build is used to execute build commands.
func Build(b build, args ...string) {
	cmd := exec.Command("go", args...)
	cmd.Env = os.Environ()
	cmd.Dir = b.dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("-Environments:")
	fmt.Println("   GOVER:\t" + GOVER)
	fmt.Println("   GOROOT:\t" + GOROOT)
	fmt.Println("   GOBIN:\t" + GOBIN)
	fmt.Println("   GOPATH:\t" + GOPATH)
	fmt.Println("   GOPROXY:\t" + GOPROXY)
	fmt.Println("   GOMOD:\t" + GOMOD)
	fmt.Println("   GOOS:\t" + runtime.GOOS)
	fmt.Println("   GOARCH:\t" + runtime.GOARCH)

	fmt.Println("-Target:")
	cmd.Env = append(cmd.Env, "GOARCH="+b.arch)
	cmd.Env = append(cmd.Env, "GOOS="+b.os)

	fmt.Println("   GOOS:\t" + b.os)
	fmt.Println("   GOARCH:\t" + b.arch)

	if b.ver != "" {
		switch b.arch {
		case "arm":
			cmd.Env = append(cmd.Env, "GOARM="+b.ver)
			fmt.Println("   GOARM:\t" + b.ver)
		case "amd64":
			cmd.Env = append(cmd.Env, "GOAMD64="+b.ver)
			fmt.Println("   GOAMD64:\t" + b.ver)
		}
	}

	fmt.Println("-Working Dir:")
	fmt.Printf("   %s: %s\n\n", cmd.Dir, version)

	fmt.Println("-Exec Program:")
	fmt.Printf("   path: %s\n", cmd.Path)
	fmt.Printf("   args: %v\n", cmd.Args[1:])

	fmt.Println("\n-Exec outputs:\n-----")

	if err := cmd.Run(); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}
}
