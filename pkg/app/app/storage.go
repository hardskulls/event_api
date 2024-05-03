package app

import (
	"database/sql"
	"github.com/pressly/goose/v3"
	"io/fs"
	"time"
)

func PrepareDB(driver, url string) (*sql.DB, error) {
	db, err := sql.Open(driver, url)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(7 * time.Minute)

	return db, nil
}

func RunMigrations(db *sql.DB, root fs.FS, path string) error {
	goose.SetBaseFS(root)
	if err := goose.SetDialect("clickhouse"); err != nil {
		return err
	}
	if err := goose.Up(db, path); err != nil {
		return err
	}

	return nil
}
