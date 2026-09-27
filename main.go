package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/cjvnjde/gator/internal/config"
	"github.com/cjvnjde/gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	appConfig, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(appConfig)
	db, err := sql.Open("postgres", appConfig.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	appState := state{
		config: &appConfig,
		db:     dbQueries,
	}

	appCommands := commands{
		commands: map[string]func(*state, command) error{},
	}

	appCommands.register("login", handlerLogin)
	appCommands.register("register", handlerRegister)
	appCommands.register("reset", handlerReset)
	appCommands.register("users", handlerUsers)

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
}
