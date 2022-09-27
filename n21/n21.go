package main

import "fmt"

// набор методов, которые мы хотим использовать, но они описаны только для int_64
type int_64 interface {
	printInt64()
}

func printInt64(x int64) {
	fmt.Printf("nbr:%T\n", x)
}

// адаптируем наш тип данных в тип данных, для которого уже реализован интерфейс, который
// мы хотим использовать
func adapter(x int) int64 {
	return int64(x)
}

func main() {
	var x int = 10
	printInt64(adapter(x))
}
