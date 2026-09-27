package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cjvnjde/gator/internal/config"
)

func main() {
	appConfig, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	appState := state{
		config: &appConfig,
	}

	appCommands := commands{
		commands: map[string]func(*state, command) error{},
	}

	appCommands.register("login", handlerLogin)

	arguments := os.Args

	if len(arguments) < 2 {
		log.Fatal(fmt.Errorf("expected at least one command"))
	}

	err = appCommands.run(&appState, command{
		name: arguments[1],
		args: arguments[2:],
	})
	if err != nil {
		log.Fatal(err)
	}

	appConfig, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(appConfig)
}
