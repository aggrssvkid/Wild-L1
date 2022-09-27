package main

var justString string

/*
количество символов в строке, которую вернет "createHugeString"
должно быть больше или равно размеру среза, который мы хотим присвоить в JustString
*/

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
