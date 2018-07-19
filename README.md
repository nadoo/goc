# goc
a simple compile tool for go

## Install
```bash
go get -u github.com/nadoo/goc
```

## Usage
Change current directory to your package dir, then: `goc COMMAND [ARGS]`

- COMMAND list
    ```
    b:         build package
    bd:        build dev package(-tags=dev)
    bw:        build windows package
    bwp:       build windows dev package(-tags=dev)
    bl:        build linux package
    blp:       build linux dev package(-tags=dev)
    bm:        build mac package
    r:         release package(build without debug info)
    rw:        release windows package
    rw32:      release windows package(x86)
    rl:        release linux package
    rl32:      release linux package(x86)
    rla:       release linux package(arm)
    rla5:      release linux package(arm v5)
    rla6:      release linux package(arm v6)
    rla7:      release linux package(arm v7)
    rla8:      release linux package(arm v8)
    rlm:       release linux package(mips)
    rm:        release mac package
    rm32:      release mac package(x86)
    i:         install package(x86)
    ```

- build package
    ```bash
    goc b
    ```

- build package for linux
    ```bash
    goc bl
    ```

- install package to `GOPATH/bin`
    ```bash
    goc i
    ```