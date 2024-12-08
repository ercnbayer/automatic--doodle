package postgres

import (
	"context"
	"fmt"
	"time"

	"automatic-doodle/ent"
)

type Logger interface {
	Trace(format string, v ...any)
	Info(format string, v ...any)
	Warning(format string, v ...any)
	Error(format string, v ...any)
	Fatal(format string, v ...any)
}

type ConfigModule interface {
	GetConfig(string) (string, error)
	GetConfigInt(string) (int, error)
}

func New(
	log Logger,
	cfg ConfigModule,
) *ent.Client {
	user, err := cfg.GetConfig("POSTGRES_DB_USER")
	if err != nil {
		log.Fatal(`POSTGRES_DB_USER is failed to fetch: %w`, err)
	}

	pwd, err := cfg.GetConfig("POSTGRES_DB_PWD")
	if err != nil {
		log.Fatal(`POSTGRES_DB_PWD is failed to fetch: %w`, err)
	}

	host, err := cfg.GetConfig("POSTGRES_DB_HOST")
	if err != nil {
		log.Fatal(`POSTGRES_DB_HOST is failed to fetch: %w`, err)
	}

	dbName, err := cfg.GetConfig("POSTGRES_DB_NAME")
	if err != nil {
		log.Fatal(`POSTGRES_DB_NAME is failed to fetch: %w`, err)
	}

	port, err := cfg.GetConfigInt("POSTGRES_DB_PORT")
	if err != nil {
		log.Fatal(`POSTGRES_DB_PORT is failed to fetch: %w`, err)
	}

	uri := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		user,
		pwd,
		host,
		port,
		dbName,
	)

	sleepTime := time.Duration(10) * time.Second

	var client *ent.Client

	for i := 0; i < 10; i++ {
		client, err = ent.Open("postgres", uri)
		if err == nil {
			break
		}

		log.Warning(
			`Postgres DB connection retry %d failed with: %s`,
			i,
			err,
		)
		time.Sleep(sleepTime)
	}

	if err != nil {
		log.Fatal(
			`Postgres DB connection failed: %s`,
			err,
		)
	}

	log.Info("Postgres DB successfully connected!")

	// TODO: close this?
	// Run the auto migration tool to create all schema resources.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal("Failed creating schema resources: %v", err)
	}

	return client
}
