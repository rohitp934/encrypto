// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sql

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Message struct {
	ID         uuid.UUID
	Senderid   uuid.UUID
	Receiverid uuid.UUID
	Message    string
	Timestamp  pgtype.Timestamp
}

type User struct {
	ID          uuid.UUID
	Username    string
	Password    string
	Dateofbirth time.Time
	Publickey   pgtype.Text
}