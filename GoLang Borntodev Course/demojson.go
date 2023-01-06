package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID           int
	EmployeeName string
	Phone        string
	Email        string
}

func main() {
	data, _ := json.Marshal(&employee{101, "Sirawit Kamchoom", "0996633516", "Sirawitkc@gmail.com"})
	fmt.Println(string(data))
}
