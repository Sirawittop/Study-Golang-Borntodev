package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 1000; i++ {
		fmt.Println(from, " : ", i)
	}
}
func main() {
	go f("Hi")
	go f("Sirawit")
	time.Sleep(5 * time.Second)
}
