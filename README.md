# goc

a simple compile tool for go

## Install

```bash
go get -u github.com/nadoo/goc
```

## Usage

Change current directory to your package dir, then: `goc COMMAND [ARGS]`

- build package
    ```bash
    goc b
    ```

- release package for linux
    ```bash
    goc rl
    ```

- install package to `GOBIN` or `GOPATH/bin`
    ```bash
    goc i
    ```

- command list
    ```bash
     b: 	build package
     bd: 	build dev package(-tags=dev)
     bdr: 	build dev package(-tags=dev and -race)
     bw: 	build windows package
     bwd: 	build windows dev package(-tags=dev)
     bwdr: 	build windows dev package(-tags=dev and -race)
     bl: 	build linux package
     bld: 	build linux dev package(-tags=dev)
     bldr: 	build linux dev package(-tags=dev and -race)
     blad: 	build linux dev package(-tags=dev)(arm64/arm v8)
     bladr: build linux dev package(-tags=dev and -race)(arm64/arm v8)
     bm: 	build mac package
     bmd: 	build mac dev package(-tags=dev)
     bmdr: 	build mac dev package(-tags=dev and -race)

     rw: 	release windows package
     rw32: 	release windows package(x86)
     rl: 	release linux package
     rl32: 	release linux package(x86)
     rla: 	release linux package(arm64/arm v8)
     rla5: 	release linux package(arm v5)
     rla6: 	release linux package(arm v6)
     rla7: 	release linux package(arm v7)
     rlm: 	release linux package(mips)
     rm: 	release mac package
     rm32: 	release mac package(x86)

     r: 	run current package
     i: 	install package to `GOBIN` or `GOPATH/bin`
     c: 	clean package
    ```