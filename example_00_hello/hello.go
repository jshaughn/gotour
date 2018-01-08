package main

import (
    "fmt"
	"github.com/jshaughn/gotour/stringutil"
)

func main() {
    const hw = "hello, world"
    fmt.Printf("%s\n", hw)
    fmt.Printf("REVERSE=%s\n", stringutil.Reverse(hw))
}

