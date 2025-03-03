package main

import (
	"fmt"
	"ogos/src/models/anime"
	"ogos/src/models/manga"
	"ogos/src/server"
)

func main() {
	h := anime.Model{}
	a := manga.Model{}
	fmt.Println("", h, a)

	fmt.Println("Hola buena tarde como 'etamo")
	// db.Connect()
	server.TurnOnServer()
}
