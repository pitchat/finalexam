package database

import (
	// pq is Postgres database driver
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
//	"time"
)

//dbURL for global database connection string
var dbURL string

//InitDB initial database table customer
func InitDB() {
	dbURL = os.Getenv("DATABASE_URL")
	if len(dbURL) == 0 {
		log.Fatal("Environment variable DATABASE_URL is empty")
	}
	db, err := Connect()
	if err != nil {
		log.Fatal("Can't connect db", err.Error())
	}
	//defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS customers(
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT,
			status TEXT
	);
	`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("Can't create table fatal error", err.Error())
	}

}

var pool *sql.DB = nil

// Connect open connection to database
func Connect() (*sql.DB, error) {
	// db, err := sql.Open("postgres", dbURL)
	// return db, err
	if pool == nil {
        d, err := sql.Open("postgres", dbURL)
        if err != nil {
            return nil, err
		}
		//d.SetConnMaxLifetime(time.Minute*5);
		//d.SetMaxIdleConns(0);
		//d.SetMaxOpenConns(5);
	    pool = d

    }
	return pool, nil
}

func Close() error{
	return pool.Close()
}
