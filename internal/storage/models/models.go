package models

import (
	"github.com/oklog/ulid/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Ticket struct {
	ID      ulid.ULID          `json:"ticker_id"`
	QR      string             `json:"qr"`
	User    User               `json:"user"`
	EventID primitive.ObjectID `json:"event_id"`
}

type User struct {
	ID        primitive.ObjectID `json:"user_id"`
	Image     *string            `json:"image"`
	Firstname *string            `json:"firstname"`
	Lastname  *string            `json:"lastname"`
	EventID   primitive.ObjectID `json:"event_id"`
}

type Event struct {
	ID          primitive.ObjectID `json:"event_id"`
	Venue       string             `json:"event_venue"`
	Time        time.Time          `json:"event_time"`
	EventPrefix string             `json:"event_prefix"`
}
