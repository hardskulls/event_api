package storage

import (
	"context"
	"database/sql"
	"event_api/pkg/domain/event"
)

func GetEventsByTime(
	ctx context.Context,
	conn *sql.Conn,
	et event.Type,
	start, end event.Time,
) ([]event.Event, error) {
	stmt, err := conn.PrepareContext(ctx, `
		SELECT eventID, eventType, userID, eventTime, payload
        FROM events
        WHERE eventType = ? AND eventTime BETWEEN ? AND ?;
	`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, et, start, end)
	if err != nil {
		return nil, err
	}

	var results []event.Event

	for rows.Next() {
		var e event.Event
		err = rows.Scan(&e.EventID, &e.EventType, &e.UserID, &e.EventTime, &e.Payload)
		if err != nil {
			return results, err
		}
		results = append(results, e)
	}

	return results, nil
}
