package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	// Conexión a Supabase (PostgreSQL)
	connStr := os.Getenv("DATABASE_URL")
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	router := mux.NewRouter()

	// Configurar endpoints
	router.HandleFunc("/api/matches", GetMatches).Methods("GET")
	router.HandleFunc("/api/matches/{id}", GetMatch).Methods("GET")
	router.HandleFunc("/api/matches", CreateMatch).Methods("POST")
	router.HandleFunc("/api/matches/{id}", UpdateMatch).Methods("PUT")
	router.HandleFunc("/api/matches/{id}", DeleteMatch).Methods("DELETE")

	log.Println("Servidor iniciado en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
