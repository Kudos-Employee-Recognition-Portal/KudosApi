package models

import (
	"database/sql"
	"log"
	"time"
)

type Award struct {
	ID             int       `json:"id"`
	Region         string    `json:"region"`
	Type           string    `json:"type"`
	RecipientName  string    `json:"recipient"`
	RecipientEmail string    `json:"email"`
	CreatorID      int       `json:"creator"`
	ConferralDate  time.Time `json:"date"`
	CreationDate   time.Time `json:"created"`
}

type Awards []Award
