package storage

import (
	"context"
	"database/sql"
	"event_api/pkg/domain/event"
	"event_api/pkg/domain/user"
	//_ "github.com/ClickHouse/clickhouse-go"
)

// GetTypesWithEvents retrieves the event types that have more than the specified number of events.
func GetTypesWithEvents(ctx context.Context, conn *sql.DB, eventCount int) ([]event.Type, error) {
	stmt, err := conn.Prepare(`
		SELECT eventType 
		FROM events
		GROUP BY eventType
		HAVING COUNT(*) > ?;
    `)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(eventCount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

// GetEventsFirstDayOfMonth retrieves events that occurred on the first day of the month.
func GetEventsFirstDayOfMonth(ctx context.Context, conn *sql.DB) ([]event.Event, error) {
	stmt, err := conn.Prepare(`
		SELECT eventID, eventType, userID, eventTime, payload 
		FROM events
		WHERE toDate(eventTime) = toStartOfMonth(eventTime);`,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

// GetUsersWithEvTypes returns a list of user IDs that have more than the specified event types.
func GetUsersWithEvTypes(ctx context.Context, conn *sql.DB, eventTypes int) ([]user.ID, error) {
	stmt, err := conn.Prepare(`
		SELECT userID 
		FROM events
		GROUP BY userID
		HAVING COUNT(DISTINCT eventType) > ?;`,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(eventTypes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

// InsertTestData fills the database with test data.
func InsertTestData(ctx context.Context, conn *sql.DB, events []event.Event) error {
	stmt, err := conn.Prepare(`
		INSERT INTO events (eventID, eventType, userID, eventTime, payload) 
		VALUES (?, ?, ?, ?, ?);`,
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, e := range events {
		_, err = stmt.Exec(e.EventID, e.EventType, e.UserID, e.EventTime, e.Payload)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetEventsByTime(
	ctx context.Context,
	conn *sql.DB,
	et event.Type,
	start, end event.Time,
) ([]event.Event, error) {
	stmt, err := conn.Prepare(`
		SELECT eventID, eventType, userID, eventTime, payload
        FROM events
        WHERE eventType = ? AND eventTime BETWEEN ? AND ?;
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(et, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

func CreateEvent(ctx context.Context, conn *sql.DB, e event.UnhandledEvent) error {
	stmt, err := conn.Prepare(`
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
	defer stmt.Close()

	_, err = stmt.Exec(
		e.EventType,
		e.UserID,
		e.EventTime,
		e.Payload,
	)
	if err != nil {
		return err
	}

	return nil
}
