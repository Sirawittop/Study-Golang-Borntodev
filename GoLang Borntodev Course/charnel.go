package main

import (
	"fmt"
	"time"
)

func processone(c chan string, data string) {
	c <- data
}
func main() {
	ch := make(chan string, 1)
	go processone(ch, "Dataone")
	fmt.Println(<-ch)

	time.Sleep(5 * time.Second)
}
