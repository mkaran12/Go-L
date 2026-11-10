package main

import "fmt"

var A, B int
var f float64

func main() {
	A = 7 * 6
	B = 100 - 1
	f = 10.0 / 3

	fmt.Println(A, B, "float : %.2f", f)
}
