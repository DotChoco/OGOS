package server

import (
	"fmt"
	"net/http"
	config "ogos/src"
)

var conf config.Server = config.GetSever()
var addrs string = fmt.Sprintf("%s:%d",
	conf.Server_Host,
	conf.Server_PORT)

func TurnOnServer() {
	println("http://" + addrs)
	if err := http.ListenAndServe(addrs, nil); err != nil {
		panic(err)
	}
}

func TurnOffServer() {

}
