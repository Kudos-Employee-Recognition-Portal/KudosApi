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
}

type Users []User

type Manager struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sig  string `json:"signature"`
}

type Managers []Manager

type Admin struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Admins []Admin

func GetUsers(db *sql.DB) (Users, error) {
	rows, err := db.Query(
		"SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users Users
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (u *User) GetUser(db *sql.DB) error {
	log.Println(u.Name)
	return db.QueryRow(
		"SELECT id, name, age FROM users WHERE name = ?",
		u.Name).Scan(&u.ID, &u.Name, &u.Age)
}

func GetManagers(db *sql.DB) (Managers, error) {
	rows, err := db.Query(
		"SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var managers Managers
	for rows.Next() {
		var m Manager
		err := rows.Scan(&m.ID, &m.Name, &m.Age)
		if err != nil {
			return nil, err
		}
		managers = append(managers, m)
	}
	return managers, nil
}

func (m Manager) GetManager(db *sql.DB) error {
	return db.QueryRow(
		"SELECT id, name, age FROM users WHERE id = ?",
		m.ID).Scan(&m.ID, &m.Name, &m.Age, &m.Sig)
}

func (m *Manager) CreateManager(db *sql.DB) error {
	res, err := db.Exec(
		"INSERT INTO users (name, age) VALUES (?, ?)",
		m.Name, m.Age)
	if err != nil {
		return err
	}
	insertID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	m.ID = int(insertID)
	return nil
}

func (m *Manager) UpdateManager(db *sql.DB) error {
	_, err := db.Exec(
		"UPDATE users SET name=?, age=? WHERE id=?",
		m.Name, m.Age, m.ID)
	return err
}

func (m *Manager) DeleteManager(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM users WHERE id=?",
		m.ID)
	return err
}

func GetAdmins(db *sql.DB) (Admins, error) {
	rows, err := db.Query(
		"SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var admins Admins
	for rows.Next() {
		var a Admin
		err := rows.Scan(&a.ID, &a.Name, &a.Age)
		if err != nil {
			return nil, err
		}
		admins = append(admins, a)
	}
	return admins, nil
}

func (a *Admin) GetAdmin(db *sql.DB) error {
	return db.QueryRow(
		"SELECT id, name, age FROM users WHERE id = ?",
		a.ID).Scan(&a.ID, &a.Name, &a.Age)
}

func (a *Admin) CreateAdmin(db *sql.DB) error {
	res, err := db.Exec(
		"INSERT INTO users (name, age) VALUES (?, ?)",
		a.Name, a.Age)
	if err != nil {
		return err
	}
	insertID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	a.ID = int(insertID)
	return nil
}

func (a *Admin) UpdateAdmin(db *sql.DB) error {
	_, err := db.Exec(
		"UPDATE users SET name=?, age=? WHERE id=?",
		a.Name, a.Age, a.ID)
	return err
}

func (a *Admin) DeleteAdmin(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM users WHERE id=?",
		a.ID)
	return err
}
