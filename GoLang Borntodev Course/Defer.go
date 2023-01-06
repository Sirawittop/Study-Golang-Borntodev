package main

import "fmt"

func add(value1, value2 float64) {
	result := value1 + value2
	fmt.Println("Result : ", result)
}
func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("I : ", i)
	}
}
func deferloop() {
	for j := 0; j < 10; j++ {
		defer fmt.Println("J : ", j)
	}
}
func main() {
	fmt.Println("Welcome to calculator")
	defer fmt.Println("End")
	defer add(20, 10)
	defer add(15, 15)
	defer add(12, 12)
	//LIFO Last In First Out
	loop()
	deferloop()
}
