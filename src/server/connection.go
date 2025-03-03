package server

import (
	"fmt"
	"net/http"
	config "ogos/src"
)

var addrs string = fmt.Sprintf("%s:%d", config.Server_Host, config.Server_PORT)

func TurnOnServer() {
	println("http://" + addrs)
	if err := http.ListenAndServe(addrs, nil); err != nil {
		panic(err)
	}
}
