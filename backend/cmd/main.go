package main

import (
	_ "github.com/lib/pq"
	"github.com/scmbr/renting-app/internal/app"
)

const configsDir = "configs"

// @title Renting App API
// @version 1.0
// @description API для аренды квартир
// @host localhost:8000
// @BasePath /
func main() {
	app.Run(configsDir)
}
