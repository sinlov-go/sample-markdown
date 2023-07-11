[![golang-ci](https://github.com/sinlov-go/sample-markdown/actions/workflows/golang-ci.yml/badge.svg)](https://github.com/sinlov-go/sample-markdown/actions/workflows/golang-ci.yml)
[![license](https://img.shields.io/github/license/sinlov-go/sample-markdown)](https://github.com/sinlov-go/sample-markdown)
[![go mod version](https://img.shields.io/github/go-mod/go-version/sinlov-go/sample-markdown?label=go.mod)](https://github.com/sinlov-go/sample-markdown)
[![GoDoc](https://godoc.org/github.com/sinlov-go/sample-markdown?status.png)](https://godoc.org/github.com/sinlov-go/sample-markdown/)
[![GoReportCard](https://goreportcard.com/badge/github.com/sinlov-go/sample-markdown)](https://goreportcard.com/report/github.com/sinlov-go/sample-markdown)
[![codecov](https://codecov.io/gh/sinlov-go/sample-markdown/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov-go/sample-markdown)
[![GitHub release)](https://img.shields.io/github/v/release/sinlov-go/sample-markdown)](https://github.com/sinlov-go/sample-markdown/releases)

## for what

- this project used to github golang lib project

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/sinlov-go/sample-markdown)](https://github.com/sinlov-go/sample-markdown/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

## depends

in go mod project

```bash
# warning use private git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "https://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q https://github.com/sinlov-go/sample-markdown.git

# test depends see full version
$ go list -mod readonly -v -m -versions github.com/sinlov-go/sample-markdown
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -mod=readonly -m -versions github.com/sinlov-go/sample-markdown | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## Features

- [X] `github.com/sinlov-go/sample-markdown/sample_mk`
  - markdown nodes `sample_mk.GenerateText`
  - `sample_mk.Parse` [] string to markdown nodes
- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## env

- minimum go version: go 1.18
- change `go 1.18`, `^1.18`, `1.18.10` to new go version

### libs

| lib                                 | version |
|:------------------------------------|:--------|
| https://github.com/stretchr/testify | v1.8.4  |
| https://github.com/sebdah/goldie    | v2.5.3  |

## usage

```go
package main

import (
	"fmt"
	"github.com/sinlov-go/sample-markdown/sample_mk"
)

func foo()  {

  nodes := []sample_mk.Node{
    sample_mk.NewHeader(2, "v1.0.0 (2020-01-18)"),
    sample_mk.NewHeader(3, "Features"),
    sample_mk.NewListItem("feat: new feature"),
    sample_mk.NewHeader(3, "Bug Fixes"),
    sample_mk.NewListItem("fix: new fix"),
  }

  result := sample_mk.GenerateText(nodes)
  fmt.Printf("%s\n", result)
}
```

- more example see [doc of sample_mk](https://pkg.go.dev/github.com/sinlov-go/sample-markdown/sample_mk)

# dev

```bash
# It needs to be executed after the first use or update of dependencies.
$ make init dep
```

- test code

```bash
$ make test testBenchmark
```

add main.go file and run

```bash
# run at env dev use cmd/main.go
$ make dev
```

- ci to fast check

```bash
# check style at local
$ make style

# run ci at local
$ make ci
```

## docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# clean test build
$ make dockerTestPruneLatest

# more info see
$ make helpDocker
```
