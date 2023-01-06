package main

import "os"

func main() {

	dataone := []byte("Hello Goodmorning \n Sirawit")
	err := os.WriteFile("data.txt", dataone, 0644)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("employeeName")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	datatwo := []byte("sira\nmanee\ntanawat\n")
	os.WriteFile("employeeName", datatwo, 0644)

}
