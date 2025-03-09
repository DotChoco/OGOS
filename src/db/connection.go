package db

import (
	"database/sql"
	"fmt"
	"log"
	config "ogos/src"

	_ "github.com/go-sql-driver/mysql"
)

var conf config.Database = config.GetDataBase()

var addrs string = fmt.Sprintf("%s:%d", conf.DB_Host, conf.DB_PORT)
var DB *sql.DB
var err error

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		conf.DB_User,
		conf.DB_Pass,
		addrs,
		conf.DB_Name)

	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("db connected")
	// Ping()

	// db_err := createTable(DB)
	// if db_err != nil {
	// 	fmt.Println(db_err)
	// 	return
	// }
}

func Disconnect() {
	DB.Close()
	println("db disconnected")
	// Ping()

}

func Ping() {
	// Ping
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	} else {
		println("todo chido")
	}

}
