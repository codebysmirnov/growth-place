package main

import (
	"log"

	"growth-place/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = cfg.Validate()
	if err != nil {
		log.Fatal(err)
	}
}
