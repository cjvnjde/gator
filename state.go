package main

import (
	"github.com/cjvnjde/gator/internal/config"
	"github.com/cjvnjde/gator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}
