package structural

import (
    "fmt"
    "testing"
)

func TestDecorator(t *testing.T) {
    decorator(printName)("Alex")
}

func printName(name string) {
    fmt.Printf("name: %s\n", name)
}
