package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/dechristopher/go-api-practice/Models"
)

func main() {
	log.Println("~ Starting...")
	cfg, cfgerr := parseConfig()
	if cfgerr != nil {
		log.Println("~ Failed to load configuration")
		log.Fatal(cfgerr)
	} else {
		log.Println("~ Loaded configuration")
	}
	log.Printf("Config: {Port: %v, RHost: %v, RPort: %v, RPass, %v}", cfg.Port, cfg.Redis.Hostname, cfg.Redis.Port, cfg.Redis.Password)
}

func parseConfig() (models.Config, error) {
	cfgfile, ferr := ioutil.ReadFile("./config.json")
	if ferr != nil {
		fmt.Printf("File error: %v\n", ferr)
		return models.Config{}, ferr
	}
	var cfg models.Config
	if decodererr := json.Unmarshal(cfgfile, &cfg); decodererr != nil {
		fmt.Printf("Decoder error: %v\n", decodererr)
		return models.Config{}, decodererr
	}
	return cfg, nil

}
