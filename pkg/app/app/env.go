package app

import (
	"event_api/pkg/app/server"
	"log/slog"
	"os"
	"strconv"
)

type (
	ClickHouseURL    = string
	PathToMigrations = string
)

type envParams struct {
	port           string
	dbUrl          string
	migrationsPath string
}

func GetFromEnv(
	l *slog.Logger,
	port, dbUrl, migrationsPath string,
) (server.Port, ClickHouseURL, PathToMigrations, error) {
	params := envParams{
		port:           os.Getenv(port),
		dbUrl:          os.Getenv(dbUrl),
		migrationsPath: os.Getenv(migrationsPath),
	}
	l.Info(
		"Using env vars",
		slog.String(port, params.port),
		slog.String(dbUrl, params.dbUrl),
		slog.String(migrationsPath, params.migrationsPath),
	)

	portInt, err := strconv.Atoi(params.port)
	if err != nil {
		return -1, "", "", err
	}

	return portInt, params.dbUrl, params.migrationsPath, nil
}
