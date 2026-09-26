package main

import (
	"fmt"
	"log"

	"github.com/cjvnjde/brog_aggregator/internal/config"
)

func main() {
	gatorConfig, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	gatorConfig.SetUser("cjvnjde")

	gatorConfig, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(gatorConfig)
}
