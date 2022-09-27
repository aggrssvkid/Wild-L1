package main

import (
	"fmt"
)

func main() {
	var index int = 2
	arr := []int{1, 3, 5, 7, 10}
	//вырезаем все элементы массива до того элемента, который хотим убрать
	slice := arr[0:index]
	//объединяем новый срез со всеми элементами, находящимися после элемента, который мы убираем
	slice = append(slice, arr[index+1:]...)
	fmt.Println(slice)
}
