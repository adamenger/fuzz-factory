package main

import (
    "bytes"
    "io/ioutil"
    "log"
    "golang.org/x/image/webp"
)

func main() {
    data, err := ioutil.ReadFile("corpus/crashers/564d09e75fafdc640e0cd852a2fa8b67cba3b553.quoted")
    if err != nil {
        log.Fatal(err)
    }
    r := bytes.NewReader(data)
    _, err = webp.Decode(r)
    if err != nil {
        log.Fatal(err)
    }
}
