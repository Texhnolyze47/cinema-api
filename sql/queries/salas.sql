-- name: CreateRoom :one
INSERT INTO salas (codigo, nombre, pelicula) VALUES ($1, $2, $3) RETURNING *;

