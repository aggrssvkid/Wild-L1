package main

import (
	"fmt"
	"time"
)

func main() {
	var datas, N int
	datas = 7
	N = 10

	canal := make(chan int)
	go readGo(canal) // запустили поток для считывания данных из канала

	/*
	 закидываем данные в канал. При каждой записи в канал, спим одну секунду.
	 В целом, если хотим сделать независимый таймер, можно создать отдельный поток,
	 и в этом потоке запустить ф-ю сна на N сек. После сна, соотв. Exit
	*/
	for i := 0; i < N; i++ {
		canal <- datas + i
		time.Sleep(time.Second)
	}
}

func readGo(canal chan int) {
	var x int
	// считываем данные из канала
	for {
		x = <-canal
		fmt.Println("READ:", x)
	}
}
