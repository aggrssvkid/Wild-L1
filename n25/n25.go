package main

import (
	"fmt"
	"time"
)

func main() {
	var s time.Duration = time.Second
	sleep(s * 3)
	fmt.Println("LOL")
}

func sleep(s time.Duration) {
	<-time.After(s)
}
