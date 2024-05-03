package storage

import (
	"context"
	"database/sql"
	"event_api/pkg/domain/event"
)

// GetTypesWithEventCount retrieves the event types that have more than the specified number of events.
func GetTypesWithEventCount(ctx context.Context, conn *sql.Conn, eventCount int) ([]event.Type, error) {
	stmt, err := conn.PrepareContext(ctx, `
		SELECT eventType 
		FROM events
		GROUP BY eventType
		HAVING COUNT(*) > ?;
    `)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(eventCount)
	if err != nil {
		return nil, err
	}

	var results []event.Type

	for rows.Next() {
		var eventType event.Type
		err = rows.Scan(&eventType)
		if err != nil {
			return results, err
		}
		results = append(results, eventType)
	}

	return results, nil
}
