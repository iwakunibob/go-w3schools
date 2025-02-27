package main

import "fmt"

var h int = 22 // Global Variable  initialized
var i int      // Set to 0 if not initialized
var j = 3      //  Global may not use :=

const (
	A int = 1
	B     = 3.14
	C     = "Hi!"
)

func main() {
	var a, b = 6, "Hello"
	c, d := 7, "World!"
	fmt.Println(b, d)
	fmt.Println(a, c)
	var e, f, g int = 1, 3, 5
	fmt.Print(e, f, g) // No new line
	fmt.Println(h, i, j)
	fmt.Println(A, B, C)
	fmt.Printf("%s war %v is %v\n", d, f, j) // https://www.w3schools.com/go/go_formatting_verbs.php
}
