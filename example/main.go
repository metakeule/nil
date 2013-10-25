package main

import (
	"fmt"
	"github.com/metakeule/nil"
)

func main() {
	var x interface{}
	fmt.Println("is nil ", nil.Is(x))
	nil.MustNot(x) // panics
}
