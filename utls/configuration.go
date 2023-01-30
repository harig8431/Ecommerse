package utls

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	ConnectionString string
	ServerIp         string
	ServerPort       string
}

var Config Configuration

func LoadConfiguration() Configuration {
	//initkeys()
	file, err := os.Open("utls/Config.json")

	if err != nil {
		log.Println("Cannot Open Config File", err)
	}
	decoder := json.NewDecoder(file)

	Config = Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Println("Failed to Decode the file", err)
	}
	file.Close()
	return Config
}