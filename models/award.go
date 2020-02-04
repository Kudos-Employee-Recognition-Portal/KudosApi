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

func GetAwards(db *sql.DB) (Awards, error) {
	rows, err := db.Query(
		"SELECT id, region, type, recipient, creator, date, created FROM awards")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var awards Awards
	for rows.Next() {
		var award Award
		err := rows.Scan(
			&award.ID, &award.Region, &award.Type, &award.RecipientName,
			&award.CreatorID, &award.CreationDate, &award.ConferralDate)
		if err != nil {
			return nil, err
		}
		awards = append(awards, award)
	}
	return awards, nil
}

func (award *Award) GetAward(db *sql.DB) error {
	return db.QueryRow(
		"SELECT id, region, type, recipient, creator, date, created FROM awards WHERE id = ?",
		award.ID).Scan(
		&award.Region, &award.Type, &award.RecipientName,
		&award.CreatorID, &award.CreationDate, &award.ConferralDate)
}

func (award *Award) CreateAward(db *sql.DB) error {
	res, err := db.Exec(
		"INSERT INTO awards (region, type, recipient, creator, date, created) VALUES (?, ?, ?, ?, ?, ?)",
		&award.Region, &award.Type, &award.RecipientName,
		&award.CreatorID, &award.CreationDate, &award.ConferralDate)
	if err != nil {
		return err
	}
	insertID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	award.ID = int(insertID)
	// TODO: Insert function call here to generate LaTeX to PDF to email chain from populated award.
	return nil
}

func (award *Award) DeleteAward(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM awards WHERE id=?",
		award.ID)
	return err
}
