package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/dechristopher/go-api-practice/Models"
)

var cfgerr error

func main() {
	log.Println("~ Starting...")

	if models.Conf, cfgerr = parseConfig(); cfgerr != nil {
		log.Println("~ Failed to load configuration")
		log.Fatal(cfgerr)
	} else {
		log.Println("~ Loaded configuration")
	}

	log.Printf("Config: {Port: %v, RHost: %v, RPort: %v, RPass, %v}", models.Conf.Port, models.Conf.Redis.Hostname, models.Conf.Redis.Port, models.Conf.Redis.Password)
}

func parseConfig() (models.Config, error) {
	if cfgfile, ferr := ioutil.ReadFile("./config.json"); ferr != nil {
		fmt.Printf("File error: %v\n", ferr)
		return models.Config{}, ferr
	} else {
		var cfg models.Config
		if decodererr := json.Unmarshal(cfgfile, &cfg); decodererr != nil {
			fmt.Printf("Decoder error: %v\n", decodererr)
			return models.Config{}, decodererr
		} else {
			return cfg, nil
		}
	}
}
