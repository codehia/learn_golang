package main

// importing package fmt which is used in the program
import (
	"fmt"
)

func main() {
	// DOT NOTATION to use Println function from previously IMPORTED fmt package
	fmt.Println("Hello World, This is a package intro code")
	// Println is a VARIADIC FUNCTION which accepts any number of PARAMETERS using (...interface{}) parameter
	// func Println(a ...interface{}) (n int, err error)
	// Println returns 2 values, number of bytes returned and any write error encountered
	n, _ := fmt.Println("Meaning of life", 42, true) // _ being used to throw away the error value returned by Println since go compiler doesn't allow unused variables
	fmt.Println(n)
}
