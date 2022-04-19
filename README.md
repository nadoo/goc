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
     bw:         build windows package
     bwd:        build windows dev package(-tags=dev)
     bwdr:       build windows dev package(-tags=dev and -race)
     bl:         build linux package
     bl3:        build linux package(amdv3)
     bld:        build linux dev package(-tags=dev)
     bld3:       build linux dev package(-tags=dev and amdv3)
     bldr:       build linux dev package(-tags=dev and -race)
     blad:       build linux dev package(-tags=dev)(arm64/arm v8)
     bladr:      build linux dev package(-tags=dev and -race)(arm64/arm v8)
     bm:         build mac package
     bmd:        build mac dev package(-tags=dev)
     bmdr:       build mac dev package(-tags=dev and -race)
     bma:        build mac package(arm64)
     bmad:       build mac dev package(-tags=dev)(arm64)
     bmadr:      build mac dev package(-tags=dev and -race)(arm64)
     rw:         release windows package
     rw3:        release windows package(amd64v3)
     rw32:       release windows package(x86)
     rl:         release linux package
     rl3:        release linux package(amd64v3)
     rl32:       release linux package(x86)
     rla:        release linux package(arm64/arm v8)
     rla5:       release linux package(arm v5)
     rla6:       release linux package(arm v6)
     rla7:       release linux package(arm v7)
     rlm:        release linux package(mips)
     rlmle:      release linux package(mipsle)
     rm:         release mac package
     rm3:        release mac package(amd64v3)
     rma:        release mac package(arm64)
     r:          run current package
     i:          install package to `GOBIN` or `GOPATH/bin`
     c:          clean package
    ```