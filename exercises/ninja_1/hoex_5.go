package main

import ("fmt")

type pilot int
var x pilot
var y int

func main() {
    fmt.Println(x)
    fmt.Printf("%T\n", x)
    x = 42
    fmt.Println(x)
    // Conversion
    y = int(x)
    fmt.Println(y)
    fmt.Printf("%T\n", y)
}
