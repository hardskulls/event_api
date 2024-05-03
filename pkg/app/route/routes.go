package route

import (
	"context"
	"database/sql"
	"errors"
	"event_api/pkg/app/server"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func badRequest(c *fiber.Ctx, msg string) error {
	return c.Status(http.StatusBadRequest).SendString(msg)
}

func internalError(c *fiber.Ctx) error {
	return c.Status(http.StatusInternalServerError).SendString("Internal server error.")
}

func getConn(ctx context.Context, c *fiber.Ctx) (*sql.Conn, error) {
	config, valid := c.Locals("cfg").(server.Config)
	if !valid {
		return nil, errors.New("config not found")
	}

	return config.DB.Conn(ctx)
}
