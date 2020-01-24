package main

import (
	"database/sql"
	"errors"
)

type user struct {
	ID   int
	Name string
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	return nil, errors.New("getUsers needs work.")
}

func (user *user) getUser(db *sql.DB) error {
	return errors.New("getUser needs work.")
}

func (user *user) createUser(db *sql.DB) error {
	return errors.New("createUser needs work.")
}

func (user *user) updateUser(db *sql.DB) error {
	return errors.New("updateUser needs work.")
}

func (user *user) deleteUser(db *sql.DB) error {
	return errors.New("deleteUser needs work.")
}
