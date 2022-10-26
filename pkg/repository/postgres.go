package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const todoTable = "todo"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
