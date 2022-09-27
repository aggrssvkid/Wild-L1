package main

import (
	"fmt"
)

func main() {
	// Создаем необходимые каналы.
	c := gen(2, 3, 5)
	out := sq(c)

	// Выводим значения.
	fmt.Println("NEXT")
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
	fmt.Println(<-out)
}

// в этом потоке канал обрабатывает данные из главной горутины
// и отправляет их в следующий канал
func gen(nums ...int) chan int {
	first := make(chan int)
	go func() {
		for _, x := range nums {
			fmt.Println("Kek")
			first <- x * 2
		}
		close(first)
	}()
	fmt.Println("Soon out")
	return first
}

// принимает данные из первого канала, закидывает их во второй канал
func sq(first chan int) chan int {
	second := make(chan int)
	go func() {
		for x := range first {
			second <- x
		}
		close(second)
	}()
	return second
}
