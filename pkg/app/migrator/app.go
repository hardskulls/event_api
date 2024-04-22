package migrator

import "github.com/golang-migrate/migrate/v4"

type App struct {
	migrator *migrate.Migrate
}

func New(cfg Config) (*App, error) {
	migrationsPathOption := "file://" + cfg.migrationsPath
	var migrationsTableOption = ""
	if cfg.migrationsTable != "" {
		migrationsTableOption = "?x-migrations-table=" + cfg.migrationsTable
	}

	migrator, err := migrate.New(
		migrationsPathOption,
		cfg.dbName+"://"+cfg.storagePath+migrationsTableOption,
	)

	return &App{migrator: migrator}, err
}

func (a *App) Up() error {
	return a.migrator.Up()
}

func MustGetMigrator() *migrate.Migrate {
	cfg := MustGetConfig()
	app, err := New(cfg)
	if err != nil {
		panic(err)
	}
	return app.migrator
}
