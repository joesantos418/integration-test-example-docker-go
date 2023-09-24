package main

import (
	"database/sql"
	"fmt"
)

const (
	dbHost     = "postgres"
	dbPort     = 5432
	dbUser     = "joesantos418"
	dbPassword = "pgpass"
	dbDbname   = "api_db"
)

func saveUser(req Request) (User, error) {
	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbDbname,
	)
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		return User{}, err
	}
	defer db.Close()

	sqlStatement := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING ID;`
	id := int64(0)
	err = db.QueryRow(sqlStatement, req.Name, req.Email).Scan(&id)
	if err != nil {
		return User{}, err
	}

	return User{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
	}, nil
}
