package main

import (
	"database/sql"
	"event_api/pkg/app/routes"
	"event_api/pkg/app/server"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"net/http"
	"os"
	"strconv"
)

const createEvent = `
CREATE TABLE IF NOT EXISTS events (
    eventID Int64,
    eventType String,
    userID Int64,
    eventTime DateTime,
    payload String
) ENGINE = MergeTree
ORDER BY (eventID, eventTime)
`

func main() {
	port := os.Getenv("PORT")
	p, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("clickhouse", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Failed to open ClickHouse connection: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(createEvent)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	srv := server.New(p)

	srv.Add(logger.New())
	srv.Add(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	srv.Handle(http.MethodPost, "/api/event", routes.CreateEvent)
	srv.MustRun()
}
