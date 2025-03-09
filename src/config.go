package config

import (
	"encoding/json"
	"log"
	"os"
)

// PUBLIC

type config struct {
	// DataBase Configuration
	DB_Host string `json:"db_Host"`
	DB_PORT int    `json:"db_PORT"`
	DB_User string `json:"db_User"`
	DB_Pass string `json:"db_Pass"`
	DB_Name string `json:"db_Name"`

	// Server Configuration
	Server_Host string `json:"server_Host"`
	Server_PORT int    `json:"server_PORT"`
}

type Database struct {
	DB_Host string
	DB_PORT int
	DB_User string
	DB_Pass string
	DB_Name string
}

type Server struct {
	Server_Host string
	Server_PORT int
}

const filePath string = "D:\\Github\\Go\\OGOS\\.config"

var dataConf config = getConf()

func getConf() config {
	var dconf config = config{}
	data, err := os.ReadFile(filePath)

	if err != nil {
		println(err)
	}

	// Deserializar el contenido JSON en la variable
	if err := json.Unmarshal(data, &dconf); err != nil {
		log.Fatalf("Error al deserializar el JSON: %v", err)
	}

	return dconf
}

func GetSever() Server {
	return Server{
		Server_Host: dataConf.Server_Host,
		Server_PORT: dataConf.Server_PORT,
	}
}

func GetDataBase() Database {
	return Database{
		DB_Host: dataConf.DB_Host,
		DB_PORT: dataConf.DB_PORT,
		DB_User: dataConf.DB_User,
		DB_Pass: dataConf.DB_Pass,
		DB_Name: dataConf.DB_Name,
	}
}
