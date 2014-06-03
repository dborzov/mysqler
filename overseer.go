package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/peter")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, is_active FROM user")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var name string
		var active bool
		if err := rows.Scan(&name, &active); err != nil {
			panic(err.Error())
		}
		fmt.Println("-----------------------------------")
		fmt.Println("Name: ", name)
		fmt.Println("Active: ", active)
	}
}
