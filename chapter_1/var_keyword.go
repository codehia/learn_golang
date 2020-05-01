package main

import (
	"fmt"
)

var x = 42 // The scope of variables with var keyword is across the program and an be accessed inside functions as well
//DECLARE there is a VARIABLE with the IDENTIFIER "z"
// Creates a variable of the given type and assigns the  zero value of the specified type to the variable

var z int

func main() {

	y := 7 // The scope of short declaration is limmited to the function and can't be accessed outside the function the declaration is used in
	fmt.Println(x)
	fmt.Println(y)
	foo()
}
func foo() {
	fmt.Println(z)
	fmt.Println(x)
}
