package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var MySQLAddress = flag.String("address", "root@tcp(localhost:3306)/peter", "Database connection address")

func main() {
	flag.Parse()
	db, err := sql.Open("mysql", *MySQLAddress)
	if err != nil {
		fmt.Println("Unable to connect to that address, baby")
		fmt.Println("-----------------------------------")
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(`SELECT 
								product.id, 
								product.name,
								product.category_id,
								category.name
							FROM 
							    product,
							    category
							WHERE
							    product.category_id=category.id;`)

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var id, name, category_id, category_name string
		if err := rows.Scan(&id, &name, &category_id, &category_name); err != nil {
			panic(err.Error())
		}
		fmt.Println("-----------------------------------")
		fmt.Println("Name: ", name)
		fmt.Println("Category: ", category_name)
	}
}
