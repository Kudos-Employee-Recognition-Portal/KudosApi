package models

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type Award struct {
	ID             int            `json:"id"`
	Region         string         `json:"region"`
	Type           string         `json:"type"`
	RecipientName  string         `json:"recipientname"`
	RecipientEmail string         `json:"recipientemail"`
	CreatorID      int            `json:"creatorid"`
	Timestamp      mysql.NullTime `json:"timestamp"`
}

type Awards []Award

func GetAwards(db *sql.DB) (Awards, error) {
	rows, err := db.Query(
		"SELECT awardID, regionID, type, recipientName, recipientEmail, creatorID, timestamp FROM award")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var awards Awards
	for rows.Next() {
		var award Award
		err := rows.Scan(
			&award.ID, &award.Region, &award.Type, &award.RecipientName,
			&award.RecipientEmail, &award.CreatorID, &award.Timestamp)
		if err != nil {
			return nil, err
		}
		awards = append(awards, award)
	}
	return awards, nil
}

// TODO: Query based search functions.

func (award *Award) GetAward(db *sql.DB) error {
	return db.QueryRow(
		"SELECT regionID, type, recipientName, recipientEmail, creatorID, timestamp FROM award WHERE awardID=?",
		award.ID).Scan(
		&award.Region, &award.Type, &award.RecipientName,
		&award.RecipientEmail, &award.CreatorID, &award.Timestamp)
}

func (award *Award) CreateAward(db *sql.DB) error {
	res, err := db.Exec(
		"INSERT INTO award (regionID, type, recipientName, recipientEmail, creatorID, timestamp) VALUES (?, ?, ?, ?, ?, ?)",
		&award.Region, &award.Type, &award.RecipientName,
		&award.RecipientEmail, &award.CreatorID, &award.Timestamp)
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
		"DELETE FROM award WHERE awardID=?",
		award.ID)
	return err
}
