package main

import "fmt"

func main() {
	var res int64
	var take int64 = 110
	var offset int = 3 // и-тый бит
	var mask int64 = 1 << offset

	//применяем операцию побитовое или.
	res = take | mask
	// если число не изменилось, значит бит не изменился, значит он рваен 1
	if res == take {
		res = take ^ mask
	}
	fmt.Println(res)
}
