package app

import (
	"belajar-go-rest/helper"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var dbUrl string = "user=axata password=axataposkenari host=localhost port=5433 dbname=belajar_go_rest sslmode=disable"

func NewDB() *sql.DB {
	db, err := sql.Open("pgx", dbUrl)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}