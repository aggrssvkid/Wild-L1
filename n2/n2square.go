package main

import (
	"fmt"
	"time"
)

func main() {
	arr := [6]int{2, 4, 6, 8, 10}
	// создаем потоки. Они будут выполняться в случайном порядке
	for i, value := range arr {
		go getSquare(i, value)
	}
	// ждем завершения всех потоков, одной секундочки им вполне хватит
	time.Sleep(time.Second)
}

// входная точка для потока
func getSquare(i int, value int) {
	fmt.Println("Thread", i, ":", value*value)
}
