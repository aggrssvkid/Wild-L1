package main

import (
	"fmt"
	"sync"
)

func main() {
	var w8 sync.WaitGroup // переменная для синхронизации потоков.

	for i := 0; i < 5; i++ {
		w8.Add(1) // увеличиваем счетчик потоков. Глаынй поток будет ожидать завершения всех
		// горутин, лишь затем, вся программа завершиться
	}

	for i := 0; i < 5; i++ {
		go printSth(&w8)
	}

	w8.Wait() // функция ожидания потоков
	fmt.Print("End main routine!\n")
}

func printSth(w8 *sync.WaitGroup) {
	defer w8.Done()
	fmt.Print("kek!\n")
}
