package models

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

// Award struct reflects the award db entity after being joined to relevant tables, receiving values rather than keys.
type Award struct {
	ID             int            `json:"id"`
	Region         region         `json:"region"`
	Type           string         `json:"type"`
	RecipientName  string         `json:"recipientname"`
	RecipientEmail string         `json:"recipientemail"`
	CreatedBy      User           `json:"createdby"`
	CreatedOn      mysql.NullTime `json:"createdon"`
	// and an accessory field to make queries easier, not sure if best practice.
	QueryDates dateRange `json:"daterange"`
}

type region struct {
	ID   int    `json:"regionid"`
	Name string `json:"regionname"`
}

type dateRange struct {
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
}

type Awards []Award

func GetAllAwards(db *sql.DB) (Awards, error) {
	rows, err := db.Query(
		"SELECT a.id, r.id, r.name, a.type, a.recipientName, a.recipientEmail, a.createdOn, m.user_id, m.firstName, m.lastName " +
			"FROM award a " +
			"JOIN region r ON a.region_id = r.id " +
			"JOIN manager m ON a.createdBy = m.user_id ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var awards Awards
	for rows.Next() {
		var award Award
		err := rows.Scan(
			&award.ID, &award.Region.ID, &award.Region.Name, &award.Type, &award.RecipientName,
			&award.RecipientEmail, &award.CreatedOn,
			&award.CreatedBy.ID, &award.CreatedBy.FirstName, &award.CreatedBy.LastName)
		if err != nil {
			return nil, err
		}
		awards = append(awards, award)
	}
	return awards, nil
}

func (award *Award) QueryAwards(db *sql.DB) (Awards, error) {
	rows, err := db.Query("SELECT a.id, r.id, r.name, a.type, a.recipientName, a.recipientEmail, a.createdOn, "+
		"m.user_id, m.firstName, m.lastName "+
		"FROM award a "+
		"JOIN region r ON a.region_id = r.id "+
		"JOIN manager m ON a.createdBy = m.user_id "+
		"WHERE (a.createdOn BETWEEN ? AND ?) "+
		"AND (? IS NULL OR a.type LIKE ?) "+
		"AND (? IS NULL OR a.recipientName LIKE ?) "+
		"AND (? IS NULL OR a.recipientEmail LIKE ?) "+
		"AND (? IS NULL OR r.name LIKE ?) ",
		award.QueryDates.StartDate, award.QueryDates.EndDate,
		award.Type, "%"+award.Type+"%",
		award.RecipientName, "%"+award.RecipientName+"%",
		award.RecipientEmail, "%"+award.RecipientEmail+"%",
		award.Region.Name, "%"+award.Region.Name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var awards Awards
	for rows.Next() {
		var award Award
		err := rows.Scan(
			&award.ID, &award.Region.ID, &award.Region.Name, &award.Type, &award.RecipientName,
			&award.RecipientEmail, &award.CreatedOn,
			&award.CreatedBy.ID, &award.CreatedBy.FirstName, &award.CreatedBy.LastName)
		if err != nil {
			return nil, err
		}
		awards = append(awards, award)
	}
	return awards, nil
}

func (award *Award) GetAward(db *sql.DB) error {
	return db.QueryRow(
		"SELECT a.id, r.id, r.name, a.type, a.recipientName, a.recipientEmail, a.createdOn, "+
			"m.user_id, m.firstName, m.lastName, m.signatureURL "+
			"FROM award a "+
			"JOIN region r ON a.region_id = r.id "+
			"JOIN manager m ON a.createdBy = m.user_id "+
			"WHERE a.id = ?",
		award.ID).Scan(
		&award.ID, &award.Region.ID, &award.Region.Name, &award.Type, &award.RecipientName,
		&award.RecipientEmail, &award.CreatedOn,
		&award.CreatedBy.ID, &award.CreatedBy.FirstName, &award.CreatedBy.LastName,
		&award.CreatedBy.SigURL)
}

func (award *Award) CreateAward(db *sql.DB) error {
	res, err := db.Exec(
		"INSERT INTO award (region_id, type, recipientName, recipientEmail, createdBy) "+
			"VALUES (?, ?, ?, ?, ?)",
		&award.Region.ID, &award.Type, &award.RecipientName,
		&award.RecipientEmail, &award.CreatedBy.ID)
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
		"DELETE FROM award WHERE id = ?",
		award.ID)
	return err
}
