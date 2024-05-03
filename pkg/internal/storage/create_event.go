package storage

import (
	"context"
	"database/sql"
	"event_api/pkg/domain/event"
)

func CreateEvent(ctx context.Context, conn *sql.Conn, e event.UnhandledEvent) error {
	stmt, err := conn.PrepareContext(ctx, `
		INSERT INTO events (eventID, eventType, userID, eventTime, payload)
		SELECT
		    COALESCE(MAX(eventID), 0) + 1,
		    ?,
		    ?,
		    ?,
		    ?
		FROM events;
	`)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, e.EventType, e.UserID, e.EventTime, e.Payload)

	return err
}
