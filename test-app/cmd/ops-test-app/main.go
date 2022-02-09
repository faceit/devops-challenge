package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

type AppConfig struct {
	Port             int    `envconfig:"PORT" default:"8080"`
	PostgresHost     string `envconfig:"POSTGRESQL_HOST" default:"localhost"`
	PostgresPort     int    `envconfig:"POSTGRESQL_PORT" default:"5432"`
	PostgresUser     string `envconfig:"POSTGRESQL_USER" default:"postgres"`
	PostgresPassword string `envconfig:"POSTGRESQL_PASSWORD" default:"mysecretpassword"`
	PostgresDbName   string `envconfig:"POSTGRESQL_DBNAME" default:"postgres"`
}

func (cfg *AppConfig) GetPostgresDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDbName,
	)
}

func main() {
	if err := run(); err != nil {
		logrus.WithError(err).Fatalf("error running application")
	}
}

func run() error {
	cfg := new(AppConfig)
	if err := envconfig.Process("", cfg); err != nil {
		return fmt.Errorf("failed to process environment variables: %w", err)
	}

	// sql.Open doesn't actually attempt to connect to the database
	db, err := sql.Open("postgres", cfg.GetPostgresDSN())
	if err != nil {
		return fmt.Errorf("cannot connect to %s with error: %s", cfg.PostgresHost, err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// we need to run a ping to make sure that to connect to the database is successful.
	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to connect to postgres on startup of the application: %w", err)
	}

	// register the health endpoint
	http.HandleFunc("/health", func(res http.ResponseWriter, req *http.Request) {
		logrus.Info("got request for /health endpoint")
		// run a ping against the database
		if err := db.PingContext(req.Context()); err != nil {
			logrus.WithError(err).Error("error running ping command against postgres")
			res.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprintf(res, "DOWN\n")
			return
		}
		logrus.Info("health endpoint is ok")
		res.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(res, "UP\n")
	})

	// listen and serve
	addr := fmt.Sprintf(":%d", cfg.Port)
	logrus.Infof("attempting to start application on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		return fmt.Errorf("failed to start http server: %w", err)
	}
	return nil
}
