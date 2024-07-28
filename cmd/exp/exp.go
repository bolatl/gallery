package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	// here first argument "pgx" is just driver name to register, next argument provides info about dataSourceName for connection
	db, err := sql.Open("pgx", "host=localhost port=5432 user=baloo password=junglebook dbname=lenslocked sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users(
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS orders(
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			amount INT,
			description TEXT
		);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tables are created")
	// name, email := "New", "New@gmail.com"
	// row := db.QueryRow(`
	// 	INSERT INTO users (name, email)
	// 	VALUES($1, $2) RETURNING id;`, name, email)
	// var id int
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User created. id =", id)
	// id := 1
	// row := db.QueryRow(`
	// 	SELECT name, email
	// 	FROM users
	// 	WHERE id=$1;`, id)
	// var name, email string
	// err = row.Scan(&name, &email)
	// if err == sql.ErrNoRows {
	// 	fmt.Println("No such record")
	// }
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("User info: name = %s, email = %s\n", name, email)

	// for i := 1; i <= 5; i++ {
	// 	amount := i * 3
	// 	desc := fmt.Sprintf("Fake order #%d", i)
	// 	_, err := db.Exec(`
	// 		INSERT INTO orders(user_id, amount, description)
	// 		VALUES($1, $2, $3);`, userId, amount, desc)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// Querying multiple records
	type Order struct {
		id     int
		userID int
		amount int
		desc   string
	}
	userId := 1
	var orders []Order
	rows, err := db.Query(`SELECT id, amount, description 
		FROM orders 
		WHERE user_id=$1;`, userId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		order.userID = userId
		err = rows.Scan(&order.id, &order.amount, &order.desc)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Orders: ", orders)
}
