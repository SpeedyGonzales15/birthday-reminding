package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Login     string    `json:"login"`
	Password  string    `json:"-"`
	Birthday  time.Time `json:"birthday"`
}

func SignUp(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err := db.Exec(query, name, email)
	if err != nil {
		return err
	}
}
