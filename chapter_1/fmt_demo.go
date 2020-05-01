package main

import (
	"fmt"
)

var y = 42

func main() {
	fmt.Println(y)         // basic print, prints the value passed
	fmt.Printf("%T\n", y)  // format string %T prints type of y
	fmt.Printf("%b\n", y)  // prints y in bytes format
	fmt.Printf("%x\n", y)  // prints y in hexadecimal
	fmt.Printf("%#x\n", y) // prints y in hexadecimal with the #
	y = 911
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y) // Prints the string format to be stored in variable s
	fmt.Println(s)
	fmt.Printf("%v", y)
}
