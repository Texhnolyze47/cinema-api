-- +goose Up
CREATE TABLE Salas
(
    codigo   UUID PRIMARY KEY,
    nombre   VARCHAR(100) NOT NULL
);

-- +goose Down
DROP TABLE Salas;
