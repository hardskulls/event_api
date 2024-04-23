package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"event_api/pkg/app/server"
	"event_api/pkg/domain/event"
	"event_api/pkg/internal/storage"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func CreateEvent(app *server.Server, c *fiber.Ctx) error {
	ctx := context.Background()

	var newEv event.UnhandledEvent

	body := c.BodyRaw()
	err := json.Unmarshal(body, &newEv)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Body is not a valid JSON")
	}

	db := c.Locals("db").(*sql.DB)

	err = storage.CreateEvent(ctx, db, newEv)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			SendString(fmt.Sprintf("Failed to create event. Error: %s", err))
	}

	return c.SendString("Created event successfully.")
}
