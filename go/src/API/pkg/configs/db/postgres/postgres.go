package postgres

import (
  "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)

const (
	host     = "database_postgres"
  port     = 5432
  user     = "my_assessment"
  password = "assessment1234"
  dbname   = "assessment"
)

func PGConnect() (*sql.DB, error) {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
		
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
		return nil, err
	}

  err = db.Ping()
  if err != nil {
    panic(err)
		return nil, err
  }
	
	return db, nil
}

func PGClose(db *sql.DB) {
	db.Close()
}