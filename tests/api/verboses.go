package api

import (
	"fmt"
	"ogos/src/models/anime"
	"ogos/src/models/manga"
)

func TestGet() {
	h := anime.Model{}
	a := manga.Model{}
	fmt.Println("", h, a)
}
