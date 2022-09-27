package main

import (
	"fmt"
	"strings"
)

func main() {
	var input_str string = "dog show sun"

	//сплитим исходную строку по пробелам, благодаря чему получим массив слов
	split_str := strings.Split(input_str, " ")

	output_str := ""
	index := len(split_str) - 1

	//пробегаемся по массиву строк, начиная с конца,
	//и приписываем новую строку к уже имеющейся
	for index >= 0 {
		output_str += split_str[index]
		output_str += " "
		index--
	}
	fmt.Println(output_str)
}
