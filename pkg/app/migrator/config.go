package migrator

import (
	"os"
)

type Config struct {
	dbName          string
	storagePath     string
	migrationsPath  string
	migrationsTable string
}

func MustGetConfig() Config {
	dbName, storagePath, migrationsPath, migrationsTable :=
		os.Getenv("DB_NAME"),
		os.Getenv("STORAGE_PATH"),
		os.Getenv("MIGRATIONS_PATH"),
		os.Getenv("MIGRATIONS_TABLE")

	if dbName == "" {
		panic("[db-name is required]")
	}
	if storagePath == "" {
		panic("[storage-path is required]")
	}
	if migrationsPath == "" {
		panic("[migrations-path is required]")
	}

	return Config{
		dbName:          dbName,
		storagePath:     storagePath,
		migrationsPath:  migrationsPath,
		migrationsTable: migrationsTable,
	}
}
