package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	canal := make(chan int)

	var i int
	var num int
	var n int = 3

	go sigTake() // этот поток ждет сигнала на выход и завершает программу

	fmt.Print("Enter number of Workers!\n")
	fmt.Scan(&num) // записываем кол-во воркеров в переменную num

	if num < 1 {
		fmt.Print("Too low workers!\n")
		return
	}
	// запускаем горутины (потоки). Кол-во горутин равно кол-ву воркеров
	for i = 0; i < num; i++ {
		go worker(i, canal)
	}
	//начинаем постоянную запись в канал из главного потока
	//(каждые 2 сек, для удобства наблюдений)
	for {
		canal <- n
		time.Sleep(time.Second * 2)
	}
}

func worker(num int, c chan int) {
	var x int

	//бесконечный цикл, у котором воркер просто считывает данные из мейн-канала.
	//если данных в канале нет, то воркер "ждет своего звездого часа"
	for {
		x = <-c
		fmt.Println("Worker num ", num)
		fmt.Println("Got it:", x)
	}
}

func sigTake() {
	signalChanel := make(chan os.Signal, 1)
	// обозначаем канал для приема сигналов и какие сигналы он будет обрабатывать
	signal.Notify(signalChanel,
		syscall.SIGINT,
		syscall.SIGQUIT)
	for {
		sig := <-signalChanel
		//если получили сигнал "ctrl+c" (SIGINT в простонародии), то выходим из программы
		if sig == syscall.SIGINT {
			fmt.Println("Got ctrl+c.\nExit")
			os.Exit(0)
		}
	}
}
