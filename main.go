package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"log"
	"movies-api/internal/database"
	"net/http"
	"os"
)

type ApiConfig struct {
	DB *database.Queries
}

type Movie struct {
	Title        string `json:"nombre"`
	Calification int32  `json:"clasificacion_edad"`
}

type Room struct {
	Name     string    `json:"nombre"`
	Pelicula uuid.UUID `json:"pelicula"`
}

func main() {

	// read the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Extract the DB_URL from the .env file
	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		log.Fatal("DB_URL not found")
	}

	// Connect to the database
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	// create a instance of the connection to the database
	apiCfg := ApiConfig{
		DB: database.New(conn),
	}
	// create a new http server mux
	mux := http.NewServeMux()

	// create a new cors handler
	handler := cors.Default().Handler(mux)

	mux.HandleFunc("/movies", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Movies API"))
	})

	// add a new route to add a movie
	mux.HandleFunc("/movies/add", func(w http.ResponseWriter, r *http.Request) {
		var newMovie Movie

		// decode the request body into a new movie
		// and check for errors
		err := json.NewDecoder(r.Body).Decode(&newMovie)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		fmt.Println(newMovie)

		// add a new movie
		createdMovie, err := apiCfg.DB.CreateMovie(r.Context(), database.CreateMovieParams{
			Codigo:            uuid.New(),
			Nombre:            newMovie.Title,
			ClasificacionEdad: newMovie.Calification,
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Movie added"))

		// return the created movie
		json.NewEncoder(w).Encode(createdMovie)

	})

	// add a new route to add a room
	mux.HandleFunc("/rooms/add", func(w http.ResponseWriter, r *http.Request) {
		var newRoom Room

		// decode the request body into a new room
		// and check for errors
		err := json.NewDecoder(r.Body).Decode(&newRoom)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		fmt.Println(newRoom)

		// add a new room
		createdRoom, err := apiCfg.DB.CreateRoom(r.Context(), database.CreateRoomParams{
			Codigo:   uuid.New(),
			Nombre:   newRoom.Name,
			Pelicula: newRoom.Pelicula,
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Room added"))

		// return the created room
		json.NewEncoder(w).Encode(createdRoom)

	})

	// consulta 1
	mux.HandleFunc("/consulta/1", func(w http.ResponseWriter, r *http.Request) {
		listNameMovies, err := apiCfg.DB.SelectNombrePeliculas(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		json.NewEncoder(w).Encode(listNameMovies)
		return
	})

	// consulta 2
	mux.HandleFunc("/consulta/2", func(w http.ResponseWriter, r *http.Request) {
		listRooms, err := apiCfg.DB.SelectPeliculasClasificacion(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		json.NewEncoder(w).Encode(listRooms)
		return

	})

	// consulta 3
	mux.HandleFunc("/consulta/3", func(w http.ResponseWriter, r *http.Request) {
		listRooms, err := apiCfg.DB.SelectPeliculasSinCalificacion(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		json.NewEncoder(w).Encode(listRooms)
		return

	})

	// consulta 4
	mux.HandleFunc("/consulta/4", func(w http.ResponseWriter, r *http.Request) {
		listRooms, err := apiCfg.DB.SelectSalasSinPelicula(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		json.NewEncoder(w).Encode(listRooms)
		return

	})

	// consulta 5
	mux.HandleFunc("/consulta/5", func(w http.ResponseWriter, r *http.Request) {
		listRooms, err := apiCfg.DB.SelectInfoSalas(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		json.NewEncoder(w).Encode(listRooms)
		return
	})

	// consulta 6
	mux.HandleFunc("/consulta/6", func(w http.ResponseWriter, r *http.Request) {
		listRooms, err := apiCfg.DB.SelectInfoPeliculas(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		json.NewEncoder(w).Encode(listRooms)
		return

	})

	// consulta 7
	mux.HandleFunc("/consulta/7", func(w http.ResponseWriter, r *http.Request) {
		listRooms, err := apiCfg.DB.SelectSalasSinPelicula(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		json.NewEncoder(w).Encode(listRooms)
		return

	})

	err = http.ListenAndServe(":8080", handler)

	fmt.Println("Server running on port 8080")
	if err != nil {
		return
	}

}
