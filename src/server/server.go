package server

import (
	"fmt"
	"net/http"
	config "ogos/src"
)

var addrs string = fmt.Sprintf("%s:%d", config.Server_URL, config.Server_PORT)

func TurnOnServer() {
	println(addrs)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.DB_PORT), nil); err != nil {
		panic(err)
	}
}
