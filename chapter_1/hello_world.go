// The package “main” tells the Go compiler that the package should compile as an executable program instead of a shared library
package main

// The pacckage to do bsic functionalities like print [Will be discussed more later]
import "fmt"

// Entry point of executable program
func main() {
	fmt.Println("Hello World")
	foo()
	fmt.Println("something more")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar() // externl function calls
}

// The end of the program when the code exits main function
func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("I'm in bar")
}

// Control Flow
// (1) Sequence
// (2) loop; iterative
// (3) conditional
