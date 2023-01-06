package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID           int
	EmployeeName string
	Phone        string
	Email        string
}

func main() {
	e := employee{}
	err := json.Unmarshal([]byte(`{"ID":101,"EmployeeName":"Sirawit Kamchoom","Phone":"0996633516","Email":"Sirawit Kamchoom"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)
	fmt.Println(e.ID)
	fmt.Println(e.EmployeeName)
	fmt.Println(e.Phone)
	fmt.Println(e.Email)
}
