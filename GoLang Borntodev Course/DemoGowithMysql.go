package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func creatingTable(db *sql.DB) {
	query := `CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
		);`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func Insert(db *sql.DB) {
	var username string
	var passwerd string
	fmt.Print("Enter username : ")
	fmt.Scanf("%s\n", &username)
	fmt.Print("Enter password : ")
	fmt.Scanf("%s\n", &passwerd)
	created_at := time.Now()
	result, err := db.Exec(`INSERT INTO users (username,password,created_at)VALUE(?,?,?)`, username, passwerd, created_at)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	fmt.Print(id)

}

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

func Delete(db *sql.DB) {
	var deleteid int
	fmt.Print("Enter your ID for Delete : ")
	fmt.Scan(&deleteid)
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, deleteid)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	db, err := sql.Open("mysql", "root:42085344720062546@(localhost:3306)/coursedb")
	if err != nil {
		fmt.Println("Fail to connect")
	} else {
		fmt.Println("Connect sucessfully")
	}
	//creatingTable(db)
	//Insert(db)
	Delete(db)
}
