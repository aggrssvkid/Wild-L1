package main

import (
	"fmt"
	"math/rand"
)

func main() {
	origin := rand.Perm(10)
	//origin := []int{5,6,7,4,5,6,8,3,5,8}
	fmt.Println("origin: ", origin)
	qsort(origin, 0, len(origin))
	fmt.Println("q-sort: ", origin)
}

func qsort(p []int, l, r int) {
	if r > l {
		mid := proc(p, l, r)
		qsort(p, l, mid)
		qsort(p, mid+1, r)
	}
}

func proc(p []int, l, r int) (i int) {
	v := p[l]
	i = l
	for j := i + 1; j < r; j++ {
		if p[j] < v {
			i++
			p[i], p[j] = p[j], p[i]
		}
	}
	p[i], p[l] = p[l], p[i]
	return i
}
