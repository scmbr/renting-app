package main

import (
	_ "github.com/lib/pq"
	"github.com/scmbr/renting-app/internal/app"
	_ "github.com/scmbr/renting-app/pkg/error"
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
