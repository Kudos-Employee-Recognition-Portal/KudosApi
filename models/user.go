package models

import "database/sql"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
	Type string `json:"type"`
}

type Users []User

type Manager struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
	Sig  string `json:"signature"`
}

type Managers []Manager

type Admin struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

type Admins []Admin

func GetUsers(db *sql.DB) (Users, error) {
	rows, err := db.Query("SELECT id, name, age FROM users")
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

func GetUser(db *sql.DB, name string) (User, error) {
	var u User
	err := db.QueryRow("SELECT id, name, age FROM users WHERE name = ?", name).Scan(&u.ID, &u.Name, &u.Age, &u.Type)
	if err != nil {
		return u, err
	}
	return u, nil
}

func GetManagers(db *sql.DB) (Users, error) {
	return nil, nil
}

func GetManager(db *sql.DB, id int) (User, error) {
	var u User
	err := db.QueryRow("SELECT id, name, age FROM users WHERE id = ?", id).Scan(&u.ID, &u.Name, &u.Age, &u.Type)
	if err != nil {
		return u, err
	}
	return u, nil
}

func CreateManager(db *sql.DB) {
	return
}

func UpdateManager(db *sql.DB) {
	return
}

func DeleteManager(db *sql.DB) {
	return
}

func GetAdmins(db *sql.DB) (Users, error) {
	return nil, nil
}

func GetAdmin(db *sql.DB, name string) (User, error) {
	var u User

	return u, nil
}

func CreateAdmin(db *sql.DB) {
	return
}

func UpdateAdmin(db *sql.DB) {
	return
}

func DeleteAdmin(db *sql.DB) {
	return
}
