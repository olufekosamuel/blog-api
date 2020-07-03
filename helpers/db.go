package helpers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "funbi1989"
	dbname   = "blog"
)

func InitDB() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s"+" password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, err
}

func CreateTables() error {
	db, err := InitDB()
	defer db.Close()

	if err != nil {
		return nil
	}

	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS users(id SERIAL, email VARCHAR PRIMARY KEY,firstname VARCHAR, lastname VARCHAR, password VARCHAR, createdat TIMESTAMP);
						`)
	_, err = db.Exec(query)

	if err != nil {
		return err
	}
	return nil
}

func DropTables() error {
	db, err := InitDB()
	defer db.Close()

	if err != nil {
		return nil
	}

	query := fmt.Sprintf(`DROP TABLE IF EXISTS users;
						`)
	// DROP TABLE IF EXISTS api_keys;
	_, err = db.Exec(query)

	if err != nil {
		return err
	}
	return nil
}
