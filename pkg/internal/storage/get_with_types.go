package storage

import (
	"context"
	"database/sql"
	"event_api/pkg/domain/user"
)

// GetUsersWithEvTypes returns a list of user IDs that have more than the specified event types.
func GetUsersWithEvTypes(ctx context.Context, conn *sql.Conn, eventTypes int) ([]user.ID, error) {
	stmt, err := conn.PrepareContext(ctx, `
		SELECT userID 
		FROM events
		GROUP BY userID
		HAVING COUNT(DISTINCT eventType) > ?;
	`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(eventTypes)
	if err != nil {
		return nil, err
	}

	var results []user.ID

	for rows.Next() {
		var u user.ID
		err = rows.Scan(&u)
		if err != nil {
			return results, err
		}
		results = append(results, u)
	}

	return results, nil
}
