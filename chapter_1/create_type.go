package main

import (
	"fmt"
)

var a int

type hotdog int // custom type by the name htdog with underlying type of primitive int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
