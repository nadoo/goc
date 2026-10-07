# goc

a simple compile tool for go

## Install

```bash
go install github.com/nadoo/goc@latest
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
    b:          build package
    bd:         build dev package(-tags=dev)
    bdr:        build dev package(-tags=dev and -race)
    bw:         build windows package(amd64)
    bwd:        build windows dev package(-tags=dev)(amd64)
    bwdr:       build windows dev package(-tags=dev and -race)(amd64)
    bl:         build linux package(amd64)
    bl3:        build linux package(amd64v3)
    bld:        build linux dev package(-tags=dev)(amd64)
    bld3:       build linux dev package(-tags=dev)(amd64v3)
    bldr:       build linux dev package(-tags=dev and -race)(amd64)
    blad:       build linux dev package(-tags=dev)(arm64)
    bladr:      build linux dev package(-tags=dev and -race)(arm64)
    bm:         build mac package(arm64)
    bmd:        build mac dev package(-tags=dev)(arm64)
    bmdr:       build mac dev package(-tags=dev and -race)(arm64)
    rw:         release windows package(amd64)
    rw3:        release windows package(amd64v3)
    rw32:       release windows package(x86)
    rl:         release linux package(amd64)
    rl3:        release linux package(amd64v3)
    rl32:       release linux package(x86)
    rla:        release linux package(arm64)
    rla5:       release linux package(arm v5)
    rla6:       release linux package(arm v6)
    rla7:       release linux package(arm v7)
    rlm:        release linux package(mips)
    rlmle:      release linux package(mipsle)
    rm:         release mac package(arm64)
    r:          run current package
    i:          install package to `GOBIN` or `GOPATH/bin`
    c:          clean package
    ```
