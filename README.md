# protoiter

`protoiter` is a Go package that provides generic iterator functions for Protocol Buffers reflection, leveraging the new `iter` package introduced in Go 1.23.

## Overview

The package offers a set of utility functions to create iterators for various Protocol Buffers entities, including:

- Descriptors
- Enum Types
- Extension Types
- Files
- Message Fields
- Message Types

## Usage Example

```go
package main

import (
    "fmt"
    "iter"

    "github.com/goaux/protoiter"
    "google.golang.org/protobuf/reflect/protoreflect"
)

func main() {
    // Assuming you have a collection of descriptors
    for i, descriptor := range protoiter.Each(yourDescriptorCollection) {
        fmt.Printf("Index: %d, Descriptor: %v\n", i, descriptor)
    }
}
```
