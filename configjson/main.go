package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type configuration struct {
	Server,
	MongoDBHost,
	DBUser,
	DBPwd,
	Database string
}

var AppConfig = configuration{}

func main() {

	loadAppConfig(&AppConfig)
	fmt.Printf("%+v\n", AppConfig.DBPwd)
}

func loadAppConfig(con *configuration) {
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	//AppConfig = Configuration{}

	err = decoder.Decode(con)

	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}
