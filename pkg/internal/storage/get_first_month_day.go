package storage

import (
	"context"
	"database/sql"
	"event_api/pkg/domain/event"
)

// GetEventsFirstDayOfMonth retrieves events that occurred on the first day of the month.
func GetEventsFirstDayOfMonth(ctx context.Context, conn *sql.Conn) ([]event.Event, error) {
	stmt, err := conn.PrepareContext(ctx, `
		SELECT eventID, eventType, userID, eventTime, payload 
		FROM events
		WHERE toDate(eventTime) = toStartOfMonth(eventTime);
	`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
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
