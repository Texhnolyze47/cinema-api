// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: salas.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createRoom = `-- name: CreateRoom :one
INSERT INTO salas (codigo, nombre, pelicula) VALUES ($1, $2, $3) RETURNING codigo, nombre, pelicula
`

type CreateRoomParams struct {
	Codigo   uuid.UUID
	Nombre   string
	Pelicula uuid.UUID
}

func (q *Queries) CreateRoom(ctx context.Context, arg CreateRoomParams) (Sala, error) {
	row := q.db.QueryRowContext(ctx, createRoom, arg.Codigo, arg.Nombre, arg.Pelicula)
	var i Sala
	err := row.Scan(&i.Codigo, &i.Nombre, &i.Pelicula)
	return i, err
}
