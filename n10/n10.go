package main

import (
	"fmt"
	"sort"
)

func main() {
	var group int
	var min int
	var i int = 0
	arr := []float64{
		-25.4, -27.0, 13.0, 19.0,
		15.5, 24.5, -21.0, 32.5,
	}
	sort.Float64s(arr)
	sets := make([][]float64, 0)

	for i < len(arr) {
		//получаем группу чисел с текущим шагом
		//первое число, не входящее в текщуюю группу, является началом новой группы чисел
		group = int(arr[i]) / 10
		//запоминаем индекс первого числа группы
		min = i
		//в цикле находим индекс последнего элемента текущей группы чисел
		for int(arr[i])/10 == group {
			i++
			if i == len(arr) {
				break
			}
		}
		//выделяем нашу группу чисел и делаем срез из основного массива
		slice := arr[min:i]
		//добавляем новый срез в матрицу срезов
		sets = append(sets, slice)
		fmt.Println(sets)
		fmt.Println(i)
	}
	fmt.Println(sets)
}
