package db

import (
	"database/sql"
	"fmt"
	config "ogos/src"

	_ "github.com/go-sql-driver/mysql"
)

var addrs string = fmt.Sprintf("%s:%d", config.DB_Host, config.DB_PORT)

// TODO: Use SurrealDB into this proyect
func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		config.DB_User,
		config.DB_Pass,
		addrs,
		config.DB_Name)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	db_err := createTable(db)
	if db_err != nil {
		fmt.Println("tabla no creada:", db_err)
		return
	}

	// fmt.Println(addrs)
	// fmt.Println(dsn)
}

func createTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS users(
			id INT NOT NULL AUTO_INCREMENT,
			name VARCHAR(50),
			email VARCHAR(50),
			PRIMARY KEY (id)
		);`
	_, err := db.Exec(query)
	return err
}

func insertUser(db *sql.DB, name string, email string) (int64, error) {
	result, err := db.Exec("INSERT INTO users(name, email) VALUES (?, ?)", name, email)

	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, err
}
