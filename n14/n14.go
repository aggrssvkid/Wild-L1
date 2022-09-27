package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x string
	getSth(x)
}

func getSth(x interface{}) {
	var str string
	//с помощью reflect.TypeOf() можно определить базовый тип, если мы его не знаем
	str = reflect.TypeOf(x).String()
	fmt.Println(str)
}
