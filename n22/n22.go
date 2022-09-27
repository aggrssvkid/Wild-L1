package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	fmt.Println("Please enter num 'a'")
	fmt.Scan(&a)
	fmt.Println("Please enter num 'b")
	fmt.Scan(&b)
	if a < math.Pow(2, 20) || b < math.Pow(2, 20) {
		fmt.Println("Incorrect numbers")
		return
	}

	var res float64

	res = a + b
	if res < 0 {
		fmt.Println("Too big sum res")
	} else {
		fmt.Println("sum:", res)
	}
	fmt.Println("raznost:", a-b)
	res = a * b
	if res < 0 {
		fmt.Println("Too big multiply res")
	} else {
		fmt.Println("multiply:", res)
	}
	fmt.Println("delim:", a/b)
}
