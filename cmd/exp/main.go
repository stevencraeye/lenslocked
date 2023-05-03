package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host string 
	Port string 
	User string
	Password string 
	Database string
	SSLMode string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func main() {
	cfg := PostgresConfig{
		Host: "localhost",
		Port: "5432",
		User: "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode: "disable",
	}
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	// Create a table... 
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			amount INT,
			description TEXT
		);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tables created.")

	// Insert some data...
	// name := "Steven"
	// email := "steve@user.com"
	// row := db.QueryRow(`
	// 	INSERT INTO users (name, email)
	// 	VALUES ($1, $2) RETURN id;
	// `, name, email)
	// var id int 
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User created.")

	// Query...
	id := 1
	row := db.QueryRow(`
		SELECT name, email 
		FROM users 
		WHERE id=$1;
	`, id)
	var name, email string 
	err = row.Scan(&name, &email)
	if err != nil {
		panic(err)
	}
	fmt.Printf("user information: name=%s, email=%s", name, email)
}