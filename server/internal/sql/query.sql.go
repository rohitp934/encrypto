// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package sql

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createMessage = `-- name: CreateMessage :one
INSERT INTO messages (
  id,
  senderid,
  receiverid,
  message,
  timestamp
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
) RETURNING id, senderid, receiverid, message, timestamp
`

type CreateMessageParams struct {
	ID         uuid.UUID
	Senderid   uuid.UUID
	Receiverid uuid.UUID
	Message    string
	Timestamp  pgtype.Timestamp
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) (Message, error) {
	row := q.db.QueryRow(ctx, createMessage,
		arg.ID,
		arg.Senderid,
		arg.Receiverid,
		arg.Message,
		arg.Timestamp,
	)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.Senderid,
		&i.Receiverid,
		&i.Message,
		&i.Timestamp,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  id,
  username,
  password,
  dateofbirth,
  publickey
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
) RETURNING id, username, password, dateofbirth, publickey
`

type CreateUserParams struct {
	ID          uuid.UUID
	Username    string
	Password    string
	Dateofbirth time.Time
	Publickey   pgtype.Text
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.Username,
		arg.Password,
		arg.Dateofbirth,
		arg.Publickey,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Dateofbirth,
		&i.Publickey,
	)
	return i, err
}

const getAllChatsForUser = `-- name: GetAllChatsForUser :many
SELECT id, senderid, receiverid, message, timestamp FROM messages WHERE senderid = $1 OR receiverid = $1 ORDER BY timestamp ASC
`

func (q *Queries) GetAllChatsForUser(ctx context.Context, senderid uuid.UUID) ([]Message, error) {
	rows, err := q.db.Query(ctx, getAllChatsForUser, senderid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.Senderid,
			&i.Receiverid,
			&i.Message,
			&i.Timestamp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMessageBySenderAndReceiver = `-- name: GetMessageBySenderAndReceiver :many
SELECT id, senderid, receiverid, message, timestamp FROM messages WHERE senderid = $1 AND receiverid = $2 ORDER BY timestamp ASC
`

type GetMessageBySenderAndReceiverParams struct {
	Senderid   uuid.UUID
	Receiverid uuid.UUID
}

func (q *Queries) GetMessageBySenderAndReceiver(ctx context.Context, arg GetMessageBySenderAndReceiverParams) ([]Message, error) {
	rows, err := q.db.Query(ctx, getMessageBySenderAndReceiver, arg.Senderid, arg.Receiverid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.Senderid,
			&i.Receiverid,
			&i.Message,
			&i.Timestamp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserById = `-- name: GetUserById :one
SELECT id, username, password, dateofbirth, publickey FROM users WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Dateofbirth,
		&i.Publickey,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, password, dateofbirth, publickey FROM users WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Dateofbirth,
		&i.Publickey,
	)
	return i, err
}
