package main

import (
	"fmt"

	"github.com/uccps-samples/build-machinery-go/make/examples/multiple-binaries/pkg/version"
)

func main() {
	fmt.Print(version.String())
}
