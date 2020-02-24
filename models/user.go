package models

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

// Would be nice to break this back into two models for admins and managers if time allows.
type User struct {
	ID        int            `json:"userid"`
	FirstName string         `json:"firstname"`
	LastName  string         `json:"lastname"`
	Email     string         `json:"email"`
	Type      string         `json:"usertype"`
	Password  string         `json:"password"`
	CreatedOn mysql.NullTime `json:"createdon"`
	CreatedBy int            `json:"createdby"`
	SigURL    sql.NullString `json:"signature"`
}

// Future: implement graphQL interface.

type Users []User

// Get basic user information, primarily for login retrieval.
func (user *User) GetUserByEmail(db *sql.DB) error {
	return db.QueryRow(
		"SELECT id, type, email, password, createdOn, createdBy FROM user WHERE email = ?",
		user.Email).Scan(&user.ID, &user.Type, &user.Email, &user.Password, &user.CreatedOn,
		&user.CreatedBy)
}

// Get basic user information by id, for admin accounts.
func (user *User) GetUserById(db *sql.DB) error {
	return db.QueryRow(
		"SELECT id, type, email, createdOn, createdBy FROM user WHERE id = ?",
		user.ID).Scan(&user.ID, &user.Type, &user.Email, &user.CreatedOn, &user.CreatedBy)
}

// Method for retrieving the more detailed manager user type information.
func (user *User) GetManagerById(db *sql.DB) error {
	return db.QueryRow(
		"SELECT id, type, email, createdOn, createdBy, firstName, lastName, signatureURL "+
			"FROM user "+
			"JOIN manager m ON user.id = m.user_id "+
			"WHERE m.user_id = ?",
		user.ID).Scan(&user.ID, &user.Type, &user.Email, &user.CreatedOn,
		&user.CreatedBy, &user.FirstName, &user.LastName, &user.SigURL)
}

// Get a list of all users.
func GetAllUsers(db *sql.DB) (Users, error) {
	var users Users
	rows, err := db.Query(
		"SELECT id, type, email, createdOn, createdBy FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Type, &user.Email,
			&user.CreatedOn, &user.CreatedBy)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Get a list of all admins.
func GetAllAdmins(db *sql.DB) (Users, error) {
	rows, err := db.Query(
		"SELECT id, type, email, createdOn, createdBy FROM user JOIN admin a ON user.id = a.user_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users Users
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Type, &user.Email, &user.CreatedOn, &user.CreatedBy)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Get a list of all managers.
func GetAllManagers(db *sql.DB) (Users, error) {
	rows, err := db.Query(
		"SELECT id, type, email, createdOn, createdBy, firstName, lastName, signatureURL " +
			"FROM user " +
			"JOIN manager m ON user.id = m.user_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users Users
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Type, &user.Email, &user.CreatedOn,
			&user.CreatedBy, &user.FirstName, &user.LastName, &user.SigURL)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Create a new admin user with transaction rollback safety.
func (user *User) CreateAdmin(db *sql.DB) error {
	tx, _ := db.Begin()
	stmt1, err := tx.Prepare(
		"INSERT INTO `user` (type, email, password, createdBy) VALUE ('admin', ?, ?, ?)")
	stmt2, err := tx.Prepare(
		"INSERT INTO `admin` (user_id) VALUE (?)")
	res, err := stmt1.Exec(user.Email, user.Password, user.CreatedBy)
	if err != nil {
		tx.Rollback()
		return err
	}
	insertID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(insertID)
	res, err = stmt2.Exec(user.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

// Create a new manager user with transaction rollback safety.
func (user *User) CreateManager(db *sql.DB) error {
	tx, _ := db.Begin()
	stmt1, err := tx.Prepare(
		"INSERT INTO `user` (type, email, password, createdBy) VALUE ('manager', ?, ?, ?)")
	stmt2, err := tx.Prepare(
		"INSERT INTO `manager` (user_id, firstName, lastName, signatureURL) VALUE (?, ?, ?, ?)")
	res, err := stmt1.Exec(user.Email, user.Password, user.CreatedBy)
	if err != nil {
		tx.Rollback()
		return err
	}
	insertID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(insertID)
	res, err = stmt2.Exec(user.ID, user.FirstName, user.LastName, user.SigURL)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

// Update a user's standard info: email address and password.
func (user *User) UpdateUserInfo(db *sql.DB) error {
	_, err := db.Exec(
		"UPDATE `user` SET email=?, password=? WHERE id=?",
		user.Email, user.Password, user.ID)
	return err
}

func (user *User) UpdateManagerInfo(db *sql.DB) error {
	_, err := db.Exec(
		"UPDATE `manager` SET firstName=?, lastName=? WHERE user_id = ?",
		user.FirstName, user.LastName, user.ID)
	return err
}

// Update a manager's signature info: this is done after the user has uploaded an image, it has been stored in the
//	datastore and a url has been generated.
func (user *User) UpdateManagerSignature(db *sql.DB) error {
	_, err := db.Exec(
		"UPDATE `manager` SET signatureURL = ? WHERE user_id = ?",
		user.SigURL, user.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete a user. By foreign key cascade, the user's admin or manager row will also be deleted as well as any awards
//	where the user is listed as the creator, in order to prevent an accumulation of orphaned rows.
func (user *User) DeleteUser(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM user WHERE id = ?",
		user.ID)
	return err
}

// Function located here due to having the common receiver, user, but notably returns a list of Awards.
func (user *User) GetManagerAwards(db *sql.DB) (Awards, error) {
	rows, err := db.Query(
		"SELECT a.id, r.id, r.name, a.type, a.recipientName, a.recipientEmail, a.createdOn, "+
			"m.user_id, m.firstName, m.lastName "+
			"FROM award a "+
			"JOIN region r ON a.region_id = r.id "+
			"JOIN manager m ON a.createdBy = m.user_id "+
			"WHERE m.user_id = ?",
		user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var awards Awards
	for rows.Next() {
		var award Award
		err := rows.Scan(
			&award.ID, &award.Region.ID, &award.Region.Name, &award.Type, &award.RecipientName,
			&award.RecipientEmail, &award.CreatedOn, &award.CreatedBy.ID, &award.CreatedBy.FirstName, &award.CreatedBy.LastName)
		if err != nil {
			return nil, err
		}
		awards = append(awards, award)
	}
	return awards, nil
}
