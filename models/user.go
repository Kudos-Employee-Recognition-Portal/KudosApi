package models

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type User struct {
	ID          int            `json:"userid"`
	FirstName   string         `json:"firstname"`
	LastName    string         `json:"lastname"`
	Email       string         `json:"email"`
	Type        int            `json:"usertype"`
	Password    string         `json:"password"`
	CreatedDate mysql.NullTime `json:"createddate"`
	CreatedBy   int            `json:"createdby"`
	SigID       sql.NullInt64  `json:"signature"`
}

type Users []User

func GetUsers(db *sql.DB) (Users, error) {
	rows, err := db.Query(
		"SELECT userID, firstName, lastName, email, userTypeID, password, createdOn, createdBy, signatureID FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users Users
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Type, &user.Password, &user.CreatedDate,
			&user.CreatedBy, &user.SigID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUsersByType(db *sql.DB, usertype int) (Users, error) {
	rows, err := db.Query(
		"SELECT userID, firstName, lastName, email, userTypeID, password, createdOn, createdBy, signatureID FROM user WHERE userTypeID = ?", usertype)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users Users
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Type, &user.Password, &user.CreatedDate,
			&user.CreatedBy, &user.SigID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (user *User) GetUser(db *sql.DB) error {
	return db.QueryRow(
		"SELECT userID, firstName, lastName, email, userTypeID, password, createdOn, createdBy, signatureID FROM user WHERE email = ?",
		user.Email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Type, &user.Password, &user.CreatedDate,
		&user.CreatedBy, &user.SigID)
}

func (user *User) GetUserByType(db *sql.DB, usertype int) error {
	return db.QueryRow(
		"SELECT userID, firstName, lastName, email, userTypeID, password, createdOn, createdBy, signatureID FROM user WHERE userID = ? AND userTypeID=?",
		user.ID, usertype).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Type, &user.Password, &user.CreatedDate,
		&user.CreatedBy, &user.SigID)
}

func (user *User) CreateAdmin(db *sql.DB) error {
	res, err := db.Exec(
		"INSERT INTO user (email, userTypeID, password, createdBy) VALUES (?, ?, ?, ?)",
		user.Email, 1, user.Password, user.CreatedBy)
	if err != nil {
		return err
	}
	insertID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(insertID)
	return nil
}

func (user *User) UpdateAdmin(db *sql.DB) error {
	_, err := db.Exec(
		"UPDATE user SET email=?, password=? WHERE userID=?",
		user.Email, user.Password, user.ID)
	return err
}

func (user *User) CreateManager(db *sql.DB) error {
	res, err := db.Exec(
		"INSERT INTO user (firstName, lastName, email, userTypeID, password, createdBy, signatureID) VALUES (?, ?, ?, ?, ?, ?, ?)",
		user.FirstName, user.LastName, user.Email, 2, user.Password, user.CreatedBy, user.SigID)
	if err != nil {
		return err
	}
	insertID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(insertID)
	return nil
}

func (user *User) UpdateManager(db *sql.DB) error {
	_, err := db.Exec(
		"UPDATE user SET firstName=?, lastName=?, email=?, password=? WHERE userID=?",
		user.FirstName, user.LastName, user.Email, user.Password)
	return err
}

func (user *User) DeleteUser(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM users WHERE userID=?",
		user.ID)
	return err
}

// TODO: implement after awards are reimplemented according to new db model.
//func (user *User) GetManagerAwards(db *sql.DB) (Awards, error) {
//	rows, err := db.Query(
//		"SELECT id, region, type, recipient, creator, date, created FROM awards WHERE creator = ?",
//		user.ID)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	var awards Awards
//	for rows.Next() {
//		var award Award
//		err := rows.Scan(
//			&award.ID, &award.Region, &award.Type, &award.RecipientName,
//			&award.CreatorID, &award.CreationDate, &award.ConferralDate)
//		if err != nil {
//			return nil, err
//		}
//		awards = append(awards, award)
//	}
//	return awards, nil
//}
