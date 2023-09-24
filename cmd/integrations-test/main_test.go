package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

const (
	dbHost     = "postgres"
	dbPort     = 5432
	dbUser     = "joesantos418"
	dbPassword = "pgpass"
	dbDbname   = "api_db"
)

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func TestInsertUser(t *testing.T) {
	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbDbname,
	)
	db, err := sql.Open("postgres", psqlconn)
	assert.Nil(t, err)
	defer db.Close()

	var actual []User
	rows, err := db.Query("SELECT * FROM users")
	assert.Nil(t, err)
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		assert.Nil(t, err)

		actual = append(actual, user)
	}
	err = rows.Err()
	assert.Nil(t, err)

	b, err := os.ReadFile("expected_outputs.json")
	assert.Nil(t, err)

	var expected []User

	err = json.Unmarshal(b, &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}
