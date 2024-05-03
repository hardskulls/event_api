package main

import (
	"event_api/pkg/app/app"
	"event_api/pkg/app/server"
	_ "github.com/ClickHouse/clickhouse-go"
	"log/slog"
)

func main() {
	l, _ := app.InitLogger()
	log := l.With(slog.String("fn", "main"))

	log.Info("Getting env vars...")
	port, dbUrl, migrationsPath, err := app.GetFromEnv(
		log, "PORT", "CLICKHOUSE_URL", "PATH_TO_MIGRATIONS",
	)
	if err != nil {
		log.Error("Failed to get one or more env variables", slog.String("error", err.Error()))
		return
	}

	log.Info("Connecting to database...")
	db, err := app.PrepareDB("clickhouse", dbUrl)
	if err != nil {
		log.Error("Failed to open the database", slog.String("error", err.Error()))
		return
	}

	log.Info("Running migrations...")
	if err = app.RunMigrations(db, nil, migrationsPath); err != nil {
		log.Error("Failed to run migrations", slog.String("error", err.Error()))
		return
	}

	log.Info("Initializing server...")
	srv := server.New(port)

	app.AddMiddleware(srv, server.Config{DB: db})

	app.AddRoutes(srv)

	log.Info("Starting server.")
	srv.MustRun()
}
