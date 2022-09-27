package main

import (
	"fmt"
)

func main() {
	arr := [6]int{2, 4, 6, 8, 10, 2}
	canal := make(chan int, len(arr)) // канал для обмена данными между потоками
	var sum int = 0

	// создаем для каждого числа из массива поток, который вычислит квадрат этого числа
	for _, value := range arr {
		go getSquare(value, canal)
	}
	//квадрат числа посылается в канал, каждым созданным потоком, а главный поток
	//считывает данные из канала и суммирует их
	for range arr {
		sum += <-canal
	}
	fmt.Println(sum)
}

func getSquare(value int, canal chan int) {
	x := value * value
	canal <- x
}
