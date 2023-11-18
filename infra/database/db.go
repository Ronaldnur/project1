package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "todos"
	dialect  = "postgres"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	db, err = sql.Open(dialect, psqlInfo)

	if err != nil {
		log.Panic("Data Koneksi yang dimasukkan salah", err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic("koneksi ke database gagal", err)
	}
}

func handleTableCreation() {
	dataTable := `
	CREATE TABLE IF NOT EXISTS "data" (
		id SERIAL PRIMARY KEY,
		title VARCHAR(250) NOT NULL,
		completed BOOLEAN
	);
	`

	createTableQuery := fmt.Sprintf("%s", dataTable)

	_, err = db.Exec(createTableQuery)

	if err != nil {
		log.Panic("terjadi kesalahan saat membuat table", err)
	}
}

func InitiazeDatabase() {
	handleDatabaseConnection()
	handleTableCreation()
}

func GetDatabaseInstance() *sql.DB {
	return db
}
