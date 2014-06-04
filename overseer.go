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
								category.name,
								vendor.name,
								vendor_inventory.regular_price
							FROM 
							    product,
							    category,
							    vendor_inventory, 
							    vendor
							WHERE
							    product.category_id=category.id 
						      AND
							    vendor_inventory.product_id=product.id
						      AND
						        vendor_inventory.vendor_id=vendor.id
							    ;`)

	if err != nil {
		panic(err.Error())
	}

	Words := make([]string, 0)
	Values := make([]interface{}, 0)

	for rows.Next() {
		var id, name, category_id, category_name, vendor, price string
		if err := rows.Scan(&id, &name, &category_id, &category_name, &vendor, &price); err != nil {
			panic(err.Error())
		}
		fmt.Println("-----------------------------------")
		fmt.Println("Name: ", name)
		fmt.Println("Category: ", category_name)
		keyword := vendor + name
		fmt.Println("Keyword: ", keyword)
		Words = append(Words, keyword)
		Values = append(Values, 10)
	}
}
