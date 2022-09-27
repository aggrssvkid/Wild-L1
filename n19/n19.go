package main

import "fmt"

func main() {
	var str string = "КОМНАТА"
	var reverse string = ""

	/*
		приводим каждый элемент нашей строки к типу rune, чтобы не потерять данные,
		когда будем применять операцию индексации string[i]. Эта операция возвращает тип byte,
		а этого типа не хватает, чтобы описать все символы unicode. (только на таблицу ascii)
	*/
	runes := []rune(str)
	var last int = len(runes) - 1

	for last >= 0 {
		reverse += string(runes[last])
		last--
	}
	fmt.Println(reverse)
}
