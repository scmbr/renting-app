package main

import (
	_ "github.com/lib/pq"
	"github.com/scmbr/renting-app/internal/app"
)

const configsDir = "configs"

func main() {
	app.Run(configsDir)
}
