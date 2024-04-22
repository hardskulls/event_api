package main

import (
	"database/sql"
	"event_api/pkg/app/api/routes"
	"event_api/pkg/app/api/server"
	"event_api/pkg/app/migrator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := os.Getenv("PORT")
	p, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("clickhouse", "tcp://localhost:9000?debug=true")
	if err != nil {
		log.Fatalf("Failed to open ClickHouse connection: %v", err)
	}
	defer db.Close()

	migrateApp := migrator.MustGetMigrator()
	err = migrateApp.Up()
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	srv := server.New(p)

	srv.Add(logger.New())
	srv.Add(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	srv.Handle(http.MethodPost, "/api/events", routes.CreateEvent)
	srv.MustRun()
}
