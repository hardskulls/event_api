package storage

import (
	"context"
	"database/sql"
	"event_api/pkg/domain/event"
	//_ "github.com/ClickHouse/clickhouse-go"
)

// InsertTestData fills the database with test data.
func InsertTestData(ctx context.Context, conn *sql.Conn, events []event.Event) error {
	stmt, err := conn.PrepareContext(ctx, `
		INSERT INTO events (eventID, eventType, userID, eventTime, payload) 
		VALUES (?, ?, ?, ?, ?);
	`)
	if err != nil {
		return err
	}

	for _, e := range events {
		_, err = stmt.Exec(e.EventID, e.EventType, e.UserID, e.EventTime, e.Payload)
		if err != nil {
			return err
		}
	}

	return nil
}
