package main

import (
	"fmt"
)

func main() {

	var a int = 7
	var b int = 5

	fmt.Println("a =", a, "b =", b)
	a = a - b
	b = b + a
	a = b - a
	fmt.Println("Change:", "a =", a, "b =", b)
}
