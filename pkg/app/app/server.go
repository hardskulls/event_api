package app

import (
	"event_api/pkg/app/route"
	"event_api/pkg/app/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"net/http"
)

const (
	API = "/api"
	V1  = "/v1"
)
const (
	Event  = "/event"
	Health = "/health"
)
const (
	Create = "/create"
)

func AddMiddleware(srv *server.Server, cfg server.Config) {
	srv.Add(logger.New())
	srv.Add(func(c *fiber.Ctx) error {
		c.Locals("cfg", cfg)
		return c.Next()
	})
}

func AddRoutes(srv *server.Server) {
	// Health check.
	srv.Handle(http.MethodGet, API+V1+Health+"/srv", route.ServerHealthCheck)
	srv.Handle(http.MethodGet, API+V1+Health+"/db", route.DatabaseHealthCheck)

	// Event routes.
	srv.Handle(http.MethodPost, API+V1+Event+Create, route.CreateEvent)
}
