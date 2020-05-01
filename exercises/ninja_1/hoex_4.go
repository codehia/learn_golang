package main

import (
	"fmt"
)

type pilot int

var x pilot

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
