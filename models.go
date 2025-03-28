package main

type Partido struct {
    ID        int    `json:"id"`
    Local     string `json:"homeTeam"`   
    Visitante string `json:"awayTeam"`   
    Fecha     string `json:"matchDate"`  
}
