package main

import (
	_ "github.com/lib/pq"
	"github.com/vasya/renting-app/internal/app"
)

func main() {
	app.Run()
}
