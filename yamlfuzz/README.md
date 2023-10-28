# yamlfuzz

this is a simple attempt at fuzzing the Unmarshal function in the gopkg.in/yaml.v3 library

## Setup
```
go get github.com/dvyukov/go-fuzz/go-fuzz
go get github.com/dvyukov/go-fuzz/go-fuzz-build
```

## Running

```
$GOPATH/bin/go-fuzz-build -preserve reflect,crypto/internal/bigmod
$GOPATH/bin/go-fuzz -bin=./yamlfuzz-fuzz.zip -workdir=workdir -procs=8
```
