package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func query(db *sql.DB) {
	var (
		id         int
		coursename string
		price      float64
		teacher    string
	)
	var inputid int
	fmt.Scanf("%d", &inputid)
	query := "SELECT id,coursename,price,teacher FROM Course WHERE id = ?"
	if err := db.QueryRow(query, inputid).Scan(&id, &coursename, &price, &teacher); err != nil {
		log.Fatal(err)
	}
	fmt.Println(id, coursename, price, teacher)
}

func main() {
	db, err := sql.Open("mysql", "root:42085344720062546@(localhost:3306)/coursedb")
	if err != nil {
		fmt.Println("Fail to connect")
	} else {
		fmt.Println("Connect sucessfully")
	}
	query(db)

}
