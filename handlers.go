package main

import (
    "encoding/json"
    "net/http"    
    "github.com/gorilla/mux"
)

func GetMatches(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, local, visitante, fecha FROM partidos")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var matches []Partido
    for rows.Next() {
        var p Partido
        if err := rows.Scan(&p.ID, &p.Local, &p.Visitante, &p.Fecha); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        matches = append(matches, p)
    }
    json.NewEncoder(w).Encode(matches)
}

func GetMatch(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    var p Partido
    err := db.QueryRow("SELECT id, local, visitante, fecha FROM partidos WHERE id = $1", vars["id"]).Scan(&p.ID, &p.Local, &p.Visitante, &p.Fecha)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(p)
}

func CreateMatch(w http.ResponseWriter, r *http.Request) {
    var p Partido
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    err := db.QueryRow(
        "INSERT INTO partidos (local, visitante, fecha) VALUES ($1, $2, $3) RETURNING id",
        p.Local, p.Visitante, p.Fecha,
    ).Scan(&p.ID)
    
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(p)
}

func UpdateMatch(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    var p Partido
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err := db.Exec(
        "UPDATE partidos SET local = $1, visitante = $2, fecha = $3 WHERE id = $4",
        p.Local, p.Visitante, p.Fecha, id,
    )
    
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(p)
}

func DeleteMatch(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    _, err := db.Exec(
        "DELETE FROM partidos WHERE id = $1",
        id,
    )
    
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusNoContent)
}