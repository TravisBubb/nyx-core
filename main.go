package main

import (
    nyx "github.com/TravisBubb/nyx-core/cmd/nyx"
    _ "github.com/TravisBubb/nyx-core/cmd/nyx/openapi"
    _ "github.com/TravisBubb/nyx-core/cmd/nyx/generate"
    _ "github.com/TravisBubb/nyx-core/cmd/nyx/run"
)

func main() {
    nyx.Execute()
}
