package route

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

func ServerHealthCheck(c *fiber.Ctx) error {
	return c.SendString("Service is live!")
}

func DatabaseHealthCheck(c *fiber.Ctx) error {
	ctx := context.Background()

	conn, err := getConn(ctx, c)
	if err != nil {
		_ = c.Status(500).SendString(err.Error())
		return err
	}

	if err = conn.PingContext(ctx); err != nil {
		_ = badRequest(c, "DB ping error")
		return err
	}

	return c.SendString("DB is live!")
}
