package models

import (
	"database/sql"
	"log"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Type string `json:"type"`
	Sig  string `json:"signature"`
}

type Users []User

func GetUsers(db *sql.DB) (Users, error) {
	rows, err := db.Query(
		"SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users Users
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Nonstandard receiver names used to reflect expected data model.
func (user *User) GetUser(db *sql.DB) error {
	log.Println(user.Name)
	return db.QueryRow(
		"SELECT id, name, age FROM users WHERE name = ?",
		user.Name).Scan(&user.ID, &user.Name, &user.Age)
}

func GetManagers(db *sql.DB) (Users, error) {
	rows, err := db.Query(
		"SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var managers Users
	for rows.Next() {
		var manager User
		err := rows.Scan(&manager.ID, &manager.Name, &manager.Age)
		if err != nil {
			return nil, err
		}
		managers = append(managers, manager)
	}
	return managers, nil
}

func (manager *User) GetManager(db *sql.DB) error {
	return db.QueryRow(
		"SELECT id, name, age FROM users WHERE id = ?",
		manager.ID).Scan(&manager.ID, &manager.Name, &manager.Age, &manager.Sig)
}

func (manager *User) CreateManager(db *sql.DB) error {
	res, err := db.Exec(
		"INSERT INTO users (name, age) VALUES (?, ?)",
		manager.Name, manager.Age)
	if err != nil {
		return err
	}
	insertID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	manager.ID = int(insertID)
	return nil
}

func (manager *User) UpdateManager(db *sql.DB) error {
	_, err := db.Exec(
		"UPDATE users SET name=?, age=? WHERE id=?",
		manager.Name, manager.Age, manager.ID)
	return err
}

func (manager *User) DeleteManager(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM users WHERE id=?",
		manager.ID)
	return err
}

func GetAdmins(db *sql.DB) (Users, error) {
	rows, err := db.Query(
		"SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var admins Users
	for rows.Next() {
		var admin User
		err := rows.Scan(&admin.ID, &admin.Name, &admin.Age)
		if err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}
	return admins, nil
}

func (admin *User) GetAdmin(db *sql.DB) error {
	return db.QueryRow(
		"SELECT id, name, age FROM users WHERE id = ?",
		admin.ID).Scan(&admin.ID, &admin.Name, &admin.Age)
}

func (admin *User) CreateAdmin(db *sql.DB) error {
	res, err := db.Exec(
		"INSERT INTO users (name, age) VALUES (?, ?)",
		admin.Name, admin.Age)
	if err != nil {
		return err
	}
	insertID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	admin.ID = int(insertID)
	return nil
}

func (admin *User) UpdateAdmin(db *sql.DB) error {
	_, err := db.Exec(
		"UPDATE users SET name=?, age=? WHERE id=?",
		admin.Name, admin.Age, admin.ID)
	return err
}

func (admin *User) DeleteAdmin(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM users WHERE id=?",
		admin.ID)
	return err
}
