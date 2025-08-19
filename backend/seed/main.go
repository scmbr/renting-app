package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/scmbr/renting-app/internal/config"
	"github.com/scmbr/renting-app/internal/infrastructure/db/postgres"
	"github.com/scmbr/renting-app/seed/seed"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func main() {
	target := flag.String("target", "all", "What to seed: users, apartments, listings, all")
	city := flag.String("city", "", "City name for seeding apartments")
	flag.Parse()

	db, cfg := InitDependencies()

	switch *target {
	case "users":
		seed.SeedUsers(db, cfg)
	case "apartments":
		seed.SeedApartments(db, cfg, *city)
	case "adverts":
		seed.SeedAdverts(db, cfg)
	case "all":
		seed.SeedUsers(db, cfg)
		seed.SeedApartments(db, cfg, *city)
		seed.SeedAdverts(db, cfg)
	default:
		log.Fatalf("unknown seed target: %s", *target)
	}
}

func InitDependencies() (*gorm.DB, *config.Config) {
	if err := godotenv.Load("../.env"); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	cfg, err := config.Init("../configs")
	if err != nil {
		log.Fatalf("failed to init config: %s", err)
	}
	fmt.Printf("Loaded config: %+v\n", cfg.Postgres)
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Username: cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.Name,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err)
	}

	return db, cfg
}
