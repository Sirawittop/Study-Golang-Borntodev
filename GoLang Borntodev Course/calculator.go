package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Printf("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		input = strings.TrimRight(input, "\r\n")
		message, err := fmt.Printf("%s ERR: not string must number only \n", input)
		if err != nil {
			fmt.Println(err)
		}
		panic(message)
	}
	return value
}
func add(value1, value2 float64) float64 {
	return value1 + value2
}
func subtract(value1, value2 float64) float64 {
	return value1 - value2
}
func multiply(value1, value2 float64) float64 {
	return value1 * value2
}
func divide(value1, value2 float64) float64 {
	return value1 / value2
}
func getOperator() string {
	fmt.Print("operator are +,=,*,/ ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func main() {
	value1 := getInput("value one : ")
	value2 := getInput("value two : ")
	var result float64
	switch operator := getOperator(); operator {
	case "+":
		result = add(value1, value2)
	case "-":
		result = subtract(value1, value2)
	case "*":
		result = multiply(value1, value2)
	case "/":
		result = divide(value1, value2)
	default:
		panic("invalid operator")
	}
	fmt.Printf("Ans = %v", result)
}
