// test.go
package main

import (
	"log"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize("db_use", "pass", "db")

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM user")
	a.DB.Exec("ALTER TABLE user AUTO_INCREMENT =1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS user
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    firstName VARCHAR(50) NOT NULL,
    lastName VARCHAR(50) NOT NULL,
    CreatedOn timestamp NOT NULL DEFAULT current_timestamp
)`
