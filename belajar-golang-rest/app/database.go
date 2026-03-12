package app

import (
	"belajar-go-rest/helper"
	"context"
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
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

func NewPgxPool() *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(dbUrl)
	helper.PanicIfError(err)

	config.MaxConns = 25
	config.MinConns = 5
	config.MaxConnLifetime = 30 * time.Minute
	config.MaxConnIdleTime = 1 * time.Minute
	config.HealthCheckPeriod = 30 * time.Second

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	helper.PanicIfError(err)

	return pool
}
