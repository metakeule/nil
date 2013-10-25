package main

import (
	"fmt"
	. "github.com/metakeule/nil"
)

func main() {
	var x interface{}
	fmt.Println("is nil ", IsNil(x))
	Nø(x) // panics
}
