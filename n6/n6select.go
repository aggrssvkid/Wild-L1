package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	done := make(chan bool)  // канал для остановки программы
	msg := make(chan string) // канал для обмена данными между потоками

	go get_line(done, msg)

	for {
		select {
		// обрабатываем данные со всех каналов
		case <-msg:
			fmt.Print("Got msg\n")
		case <-done: //как только поступит что-то в канал "done", завершаем прогу
			fmt.Print("Exit\n")
			os.Exit(0)
		}
	}
}

func get_line(done chan bool, msg chan string) {
	var line string
	for {
		fmt.Print("Enter anything\n")
		fmt.Scan(&line)
		if line == "quit" {
			done <- true
			time.Sleep(time.Second)
		} else {
			msg <- line
			time.Sleep(time.Second)
		}

	}
}
