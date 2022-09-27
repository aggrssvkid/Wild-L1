package main

import "fmt"

func main() {
	var find int = 0
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	new := make([]string, 0)

	/*
		Проходимся по элементам базового массива и, если этого элемента нет в
		Новом, создаваемом нами массиве, добавляем туда этот элемент
	*/
	for _, arr_value := range arr {
		for _, new_value := range new {
			if arr_value == new_value {
				find = 1
				break
			}
		}
		if find == 1 {
			find = 0
		} else {
			new = append(new, arr_value)
		}
	}
	fmt.Println(new)
}
