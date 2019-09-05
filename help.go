// ALL possible goos and goarch
// https://github.com/golang/go/blob/master/src/go/build/syslist.go
// const goosList = "android darwin dragonfly freebsd linux nacl netbsd openbsd plan9 solaris windows zos "
// const goarchList = "386 amd64 amd64p32 arm armbe arm64 arm64be ppc64 ppc64le mips mipsle mips64 mips64le mips64p32 mips64p32le ppc s390 s390x sparc sparc64 "

// for arm:
// https://github.com/golang/go/wiki/GoArm
//

package main

import "fmt"

func showHelp() {
	fmt.Println()
	fmt.Println("-")
	fmt.Println("Usage: goc COMMAND [ARGS]")
	fmt.Println()
	fmt.Println("COMMAND")
	fmt.Println("     b: \tbuild package")
	fmt.Println("     bd: \tbuild dev package(-tags=dev)")
	fmt.Println("     bdr: \tbuild dev package(-tags=dev and -race)")
	fmt.Println("     bw: \tbuild windows package")
	fmt.Println("     bwd: \tbuild windows dev package(-tags=dev)")
	fmt.Println("     bwdr: \tbuild windows dev package(-tags=dev and -race)")
	fmt.Println("     bl: \tbuild linux package")
	fmt.Println("     bld: \tbuild linux dev package(-tags=dev)")
	fmt.Println("     bldr: \tbuild linux dev package(-tags=dev and -race)")
	fmt.Println("     bm: \tbuild mac package")
	fmt.Println("     bmd: \tbuild mac dev package(-tags=dev)")
	fmt.Println("     bmdr: \tbuild mac dev package(-tags=dev and -race)")
	fmt.Println()
	fmt.Println("     rw: \trelease windows package")
	fmt.Println("     rw32: \trelease windows package(x86)")
	fmt.Println("     rl: \trelease linux package")
	fmt.Println("     rl32: \trelease linux package(x86)")
	fmt.Println("     rla: \trelease linux package(arm64/arm v8)")
	fmt.Println("     rla5: \trelease linux package(arm v5)")
	fmt.Println("     rla6: \trelease linux package(arm v6)")
	fmt.Println("     rla7: \trelease linux package(arm v7)")
	fmt.Println("     rlm: \trelease linux package(mips)")
	fmt.Println("     rm: \trelease mac package")
	fmt.Println("     rm32: \trelease mac package(x86)")
	fmt.Println()
	fmt.Println("     r: \trun package")
	fmt.Println("     i: \tinstall package to `GOBIN` or `GOPATH/bin`")
	fmt.Println("     c: \tclean package")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("     goc b")
	fmt.Println("       -go build in current directory")
	fmt.Println("     goc rl")
	fmt.Println("       -release linux package")
	fmt.Println()
	fmt.Println("   Ver: 1.0.0, Author: NadOo")
	fmt.Println("   Source: https://github.com/nadoo/goc")
	fmt.Println("-")
}
