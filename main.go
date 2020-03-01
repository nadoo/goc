// ALL possible goos and goarch
// https://github.com/golang/go/blob/master/src/go/build/syslist.go
// const goosList = "aix android darwin dragonfly freebsd hurd illumos js linux nacl netbsd openbsd plan9 solaris windows zos "
// const goarchList = "386 amd64 amd64p32 arm armbe arm64 arm64be ppc64 ppc64le mips mipsle mips64 mips64le mips64p32 mips64p32le ppc riscv riscv64 s390 s390x sparc sparc64 wasm "

// for arm:
// https://github.com/golang/go/wiki/GoArm
//

package main

import (
	"fmt"
	"os"
)

var version = "dev"

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
