package app

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	app_cfg "github.com/scmbr/renting-app/internal/config"
	"github.com/scmbr/renting-app/internal/handler"
	"github.com/scmbr/renting-app/internal/repository"
	"github.com/scmbr/renting-app/internal/server"
	"github.com/scmbr/renting-app/internal/service"
	"github.com/scmbr/renting-app/pkg/hash"
	"github.com/scmbr/renting-app/pkg/storage"
	"github.com/sirupsen/logrus"
)

func Run(configPath string) {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	logrus.SetFormatter(new(logrus.JSONFormatter))
	cfg, err := app_cfg.Init(configPath)
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())

		return
	}
	fmt.Print(cfg.Postgres.Password)
	fmt.Print(cfg.Postgres.Port)
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Username: cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.Name,
		SSLMode:  cfg.Postgres.SSLMode,
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db:%s", err.Error())
	}
	storageProvider, err := newStorageProvider(cfg)
	if err != nil {
		logrus.Fatalf("error initializing storage: %s", err.Error())
	}
	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)
	repos := repository.NewRepository(db)
	services := service.NewServices(service.Deps{
		Repos:           repos,
		Hasher:          hasher,
		StorageProvider: storageProvider,
		AccessTokenTTL:  cfg.Auth.JWT.AccessTokenTTL,
		RefreshTokenTTL: cfg.Auth.JWT.RefreshTokenTTL,
	})
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(cfg.HTTP.Port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func newStorageProvider(cfg *app_cfg.Config) (storage.Provider, error) {

	s3_cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion("ru-central1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.FileStorage.AccessKey,
			cfg.FileStorage.SecretKey,
			"",
		)),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(s3_cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(cfg.FileStorage.Endpoint)
		o.RequestChecksumCalculation = aws.RequestChecksumCalculationWhenRequired
		o.EndpointResolverV2 = s3.NewDefaultEndpointResolverV2()
	})
	provider := storage.NewFileStorage(client, cfg.FileStorage.Bucket, cfg.FileStorage.Endpoint, cfg.FileStorage.Website)
	return provider, nil
}
