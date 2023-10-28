# jwtfuzz

this is a simple attempt at fuzzing the Parse function in the github.com/golang-jwt/jwt library

## Setup
```
go get github.com/dvyukov/go-fuzz/go-fuzz
go get github.com/dvyukov/go-fuzz/go-fuzz-build
```

## Running

```
$GOPATH/bin/go-fuzz-build -preserve reflect,crypto/internal/bigmod
$GOPATH/bin/go-fuzz -bin=./jwtfuzz-fuzz.zip -workdir=workdir -procs=8
```
