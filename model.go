package main

import (
    "time"
)

type Match struct {
    ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
    HomeTeam  string    `json:"homeTeam" gorm:"type:varchar(255);not null"`
    AwayTeam  string    `json:"awayTeam" gorm:"type:varchar(255);not null"`
    MatchDate time.Time `json:"matchDate" gorm:"type:date;not null"`
}