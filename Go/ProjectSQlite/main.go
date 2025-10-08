package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	//pause the program
	// time.Sleep(10 * time.Second)

	db, err := sql.Open("sqlite", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected")

	defer db.Close()

	_ = db

	createTable(db)
	users := getUsers(db)
	fmt.Println(users)
}

func createTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	insertqsl := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err = db.Exec(insertqsl, "John Doe", "john.doe@example.com")
	if err != nil {
		log.Fatal(err)
	}

	insertqsl = "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err = db.Exec(insertqsl, "Jane Doe", "jane.doe@example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table created and data inserted")
}

type User struct {
	ID    int
	Name  string
	Email string
}

func getUsers(db *sql.DB) []User {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
}
