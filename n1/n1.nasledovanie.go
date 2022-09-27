package main

import (
	"fmt"
)

// Базовая структура
type Human struct {
	Name string
}

// метод для базовой структуры
func (h *Human) getROFL() {
	fmt.Println("ROFL")
}

/*
структура, которая будет наследовать все поля и методы
базовой структуры Human. То есть Human-родитель для Action
*/
type Action struct {
	Repeat int
	Human
}

// метод для дочки
func (a *Action) move() {
	fmt.Println("We rush")
}

// переопределяем метод для Action, если хотим использовать его,
// а не унаследованный метод от Human
func (a *Action) getROFL() {
	fmt.Println("KEKROFLKEK")
}

func main() {
	var Vasya Action
	Vasya.Human.getROFL()
	Vasya.getROFL()
}
