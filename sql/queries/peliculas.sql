-- name: SelectNombrePeliculas :many
SELECT nombre
FROM peliculas;

-- name: SelectPeliculasClasificacion :many
SELECT DISTINCT clasificacion_edad
FROM peliculas;

-- name: SelectPeliculasSinCalificacion :many
SELECT nombre
FROM peliculas
WHERE clasificacion_edad IS NULL OR clasificacion_edad = 0;


-- name: SelectSalasSinPelicula :many
SELECT nombre
FROM salas
WHERE NOT EXISTS (
    SELECT *
    FROM peliculas
    WHERE peliculas.codigo = salas.codigo
);

-- name: SelectInfoSalas :many
SELECT
    s.codigo,
    s.nombre,
    p.codigo AS "Código película",
    p.nombre AS "Nombre película",
    p.clasificacion_edad AS "Calificación edad"
FROM Salas s
LEFT JOIN Peliculas p ON p.codigo = s.Pelicula;

-- name: SelectInfoPeliculas :many
SELECT p.nombre as "Nombre película", p.clasificacion_edad as "Calificación edad"
FROM peliculas p
LEFT JOIN salas s ON p.codigo = s.pelicula
WHERE s.codigo IS NULL;


-- name: SelectNombrePeliculasSinSala :many
SELECT nombre
FROM peliculas
WHERE NOT EXISTS (
    SELECT *
    FROM salas
    WHERE salas.pelicula = peliculas.codigo
);


-- name: CreateMovie :one
INSERT INTO peliculas (codigo,nombre, clasificacion_edad)
VALUES ($1, $2, $3)
RETURNING *;