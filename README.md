# git module to go mod

Converts git submodules in a `vendor` directory to a `go.mod` file.

## Running

Just running the program:

```
$ gitmod2gomod --repo-path ~/go/src/go.otter.engineering/gitmoddep
```

or direct the output to the source directory:

```
$ gitmod2gomod \
 --repo-path ~/go/src/go.otter.engineering/gitmoddep > \
 ~/go/src/go.otter.engineering/gitmoddep/go.mod
```

**Generated File**

```
$ cat go.mod
module go.otter.engineering/gitmoddep

require (
  github.com/urfave/cli  934abfb2f102315b5794e15ebc7949e4ca253920
  go.uber.org/zap  67bc79d13d155c02fd008f721863ff8cc5f30659
)
```

**Running go build will update the go.mod**

```
$ GO111MODULE=on go1.11rc2 build -o ./bin/test .
go: finding github.com/urfave/cli 934abfb2f102315b5794e15ebc7949e4ca253920
go: finding go.uber.org/zap 67bc79d13d155c02fd008f721863ff8cc5f30659
```

**Final go.mod**

```
$ cat go.mod
module go.otter.engineering/gitmoddep

require (
	github.com/urfave/cli v1.20.1-0.20180821064027-934abfb2f102
	go.uber.org/zap v1.9.2-0.20180814183419-67bc79d13d15
)
```
