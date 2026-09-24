package db

import (
	"database/sql"
	"ecommerce/handlers"
	"ecommerce/models"
	"fmt"

	_ "github.com/lib/pq"
)

// var Db sql.DB
func ConnectDb() *sql.DB {
	ConnectStr := "host=localhost port=5432 user=postgres password=123456 dbname=ecommerce sslmode=disable"
	Db, err := sql.Open("postgres", ConnectStr)
	if err != nil {
		panic(err)

	}
	Err := Db.Ping()
	if Err != nil {
		fmt.Println("Database Connection Failed !")
		return nil
	}
	fmt.Println("DB connected")
	// defer Db.Close()

	sqlState := `select *from products`
	rows, err := Db.Query(sqlState)
	if err != nil {
		fmt.Println("Query error", err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var prd models.Product
		err := rows.Scan(
			&prd.ID,
			&prd.Title,
			&prd.Description,
			&prd.Price,
			&prd.ImgUrl,
		)
		if err != nil {
			fmt.Println("Scan Failed")
			return nil
		}
		handlers.Products = append(handlers.Products, prd)
		fmt.Println("ID:", prd.ID)
		fmt.Println("Title:", prd.Title)
		fmt.Println("Description:", prd.Description)
		fmt.Println("Price:", prd.Price)
		fmt.Println("Image:", prd.ImgUrl)
		fmt.Println("----------------------")
	}
	return Db
}

// func GetDb()
