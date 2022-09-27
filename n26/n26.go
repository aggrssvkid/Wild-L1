package main

import "fmt"

func main() {
	str := "eqldsakK"
	var answer bool

	answer = is_unique(str)
	fmt.Println(answer)
}

func is_unique(str string) bool {
	if len(str) == 0 {
		return true
	}

	var j int
	var i int = 0
	/*
		Пробегаемся в цикле по исходной строке. Каждый следующий элемнт сравниваем
		со всеми элементами, которые находятся после текущего элемента.
		Как только нашли совпадение, сигнализируем об этом и выходим из цикла.
		Если пробежались по всем элементам строки, то это значит, что совпадений не обнаружно
	*/
	for i < len(str)-1 {
		j = i + 1
		for j < len(str) {
			if str[i] == str[j] || is_registr(str[j], str[i]) == true {
				return false
			}
			j++
		}
		i++
	}
	return true
}

// проверяем совпадение по регистру
func is_registr(reg byte, symbol byte) bool {
	if symbol > 'a' && symbol < 'z' {
		if reg == symbol-32 {
			return true
		}
	} else if symbol > 'A' && symbol < 'Z' {
		if reg == symbol+32 {
			return true
		}
	}
	return false
}
