package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	arr := rand.Perm(20)
	sort.Ints(arr)
	fmt.Println(arr)
	// с помощью функции searchInts() осуществляем бинарный поиск по отсортированному масиву
	i := sort.SearchInts(arr, -5) //если i = len(arr), то искомый элемент больше всех элементов массива
	fmt.Println(i)
}
