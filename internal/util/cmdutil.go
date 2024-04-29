package util

import (
	"context"
	"database/sql"
	"log"
	"os"

	"go.uber.org/zap"
)

func NewLogger(service string) *zap.Logger {
	env := os.Getenv("ENV")
	logger, _ := zap.NewProduction(zap.Fields(
		zap.String("env", env),
		zap.String("service", service),
	))

	if env == "" || env == "development" {
		logger, _ = zap.NewDevelopment()
	}

	return logger
}

func NewDatabase(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./nestfile.db")
	if err != nil {
		return nil, err
	}

	sqlStmt := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );
`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}
	var userCount int
	if err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&userCount); err != nil {
		log.Fatal(err)
	}

	if userCount <= 0 {
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}
		stmt, err := tx.Prepare("insert into users(username, password) values(?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		pass,err := HashPwd("password")
		if err != nil{
			log.Fatal(err)
		}

		_, err = stmt.Exec("admin", pass)
		if err != nil {
			log.Fatal(err)
		}

		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}
	}

	err = db.PingContext(ctx)
	if err != nil {
		db.Close() // Close the database if there's an error
		return nil, err
	}

	return db, nil
}
