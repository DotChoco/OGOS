package main

import (
	"ogos/src/db"
)

func main() {
	// fmt.Println("Hola buena tarde como 'etamo")
	db.Connect()
	db.Disconnect()
}
