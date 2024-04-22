package event

import (
	"event_api/pkg/domain/user"
	"time"
)

type ID = int64
type Type = string
type Time = time.Time
type Payload = string

type Event struct {
	EventID   ID
	EventType Type
	UserID    user.ID
	EventTime Time
	Payload   Payload
}

func New(eventID int64, raw UnhandledEvent) Event {
	return Event{
		EventID:   eventID,
		EventType: raw.EventType,
		UserID:    raw.UserID,
		EventTime: raw.EventTime,
		Payload:   raw.Payload,
	}
}

type UnhandledEvent struct {
	EventType Type    `json:"eventType"`
	UserID    user.ID `json:"userId"`
	EventTime Time    `json:"eventTime"`
	Payload   Payload `json:"payload"`
}
