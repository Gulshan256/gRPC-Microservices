package main

import (
	"log"
	"time"

	"github.com/Gulshan256/go-gRPC-Microservices/account"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
)

type Config struct {
	// DatabaseURL string `envconfig:"database_url" required:"true"`
	DatabaseURL string `envconfig:"database_url"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r account.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = account.NewPostgresRepositories(cfg.DatabaseURL)
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	log.Printf("connected to database: %s", cfg.DatabaseURL)
	s := account.NewService(r)
	err = account.ListenGRPC(s, 8080)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("gRPC server started on port 8080")
}
