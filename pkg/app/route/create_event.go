package route

import (
	"context"
	"event_api/pkg/domain/convert"
	"event_api/pkg/domain/event"
	"event_api/pkg/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func CreateEvent(c *fiber.Ctx) error {
	ctx := context.Background()

	ev, err := convert.FromJSON[event.UnhandledEvent](c.BodyRaw())
	if err != nil {
		_ = internalError(c)
		return err
	}

	conn, err := getConn(ctx, c)
	if err != nil {
		_ = internalError(c)
		return err
	}

	if err = storage.CreateEvent(ctx, conn, ev); err != nil {
		_ = internalError(c)
		return err
	}

	return c.SendString("Created event successfully.")
}
