package models

import (
	"database/sql"
	"log"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

func GetAllUsers(db *sql.DB) ([]*User, error) {
	log.Println("Hit: GetUsers")
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
