package main

import (
	"fmt"
	"sync"
)

func main() {
	increment := 0
	n := 100
	wg := sync.WaitGroup{}
	//Благодаря мьютксу доступ к итерируемой переменной сможет получить только одна горутина
	//соответственно запись и чтение этой переменной будет синхронизировано
	mx := sync.Mutex{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			mx.Lock()
			increment++
			mx.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(increment)
}
