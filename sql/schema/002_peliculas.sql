-- +goose Up


CREATE TABLE Peliculas
(
    codigo           UUID PRIMARY KEY,
    nombre           VARCHAR(100) NOT NULL,
    clasificacion_edad INT NOT NULL
);

ALTER TABLE Salas ADD COLUMN Pelicula UUID NOT NULL REFERENCES Peliculas (codigo) ON DELETE CASCADE;

-- +goose Down
ALTER TABLE Salas DROP COLUMN Pelicula;
DROP TABLE Peliculas;
