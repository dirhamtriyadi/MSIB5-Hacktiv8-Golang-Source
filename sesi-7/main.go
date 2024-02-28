package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	user     = "airelljordan"
	password = "postgres"
	dialect  = "postgres"
	port     = "5432"
	dbname   = "learn-pg"
)

var (
	db  *sql.DB
	err error
)

type Product struct {
	Id    int
	Name  string
	Price int
}

func handleDatabaseConnection() {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open(dialect, psqlInfo)

	if err != nil {
		log.Panic("error occured while validating database arguments:", err.Error())
	}

	err = db.Ping()

	if err != nil {
		log.Panic("error occured while opening a connection to database:", err.Error())
	}

	fmt.Println("Successfully connected to database")
}

func createProductDataWithReturnStatement(name string, price int) {
	createProductQuery := `
		INSERT INTO "Products" 
		(name, price)
		VALUES ($1, $2)
		RETURNING id, name, price
	`

	row := db.QueryRow(createProductQuery, name, price)

	var p Product

	err := row.Scan(&p.Id, &p.Name, &p.Price)

	if err != nil {
		fmt.Println("error while scanning sql row:", err.Error())
		return
	}

	fmt.Printf("created product: %#v\n", p)

}

func createProductDataWithoutReturnStatement(name string, price int) {
	createProductQuery := `
		INSERT INTO "Products" 
		(name, price)

		VALUES ($1, $2)
	`

	// createProductQueryForbidden := fmt.Sprintf("INSERT INTO PRODUCTS (name, price) VALUES (%s, %s)", name, price)

	_, err := db.Exec(createProductQuery, name, price)

	if err != nil {
		fmt.Println("error while creating product data:", err.Error())
		return
	}

	fmt.Println("Successfully creating product data")
}

func main() {
	handleDatabaseConnection()

	createProductTable := `
		CREATE TABLE IF NOT EXISTS "Products" (
			id SERIAL,
			name varchar(255) NOT NULL,
			price int NOT NULL
		)
	`

	_, err := db.Exec(createProductTable)

	if err != nil {
		log.Panic("error while creating product table:", err.Error())
	}

	// createProductDataWithoutReturnStatement("Baseus", 15000)

	createProductDataWithReturnStatement("Iphone 15 Pro Max", 26000)
}
