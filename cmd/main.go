package main

import (
	"crud"
	"crud/pkg/handler"
	"crud/pkg/repository"
	"crud/pkg/service"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error reading from env variables %s", err.Error())
	}
	cfg := repository.Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("failed to initialize db %s", err.Error())
	}
	err = initDb(db)
	if err != nil {
		log.Fatalf("failed to initialize table %s", err.Error())
	}
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occure while running server %s", err.Error())
	}
}

func initDb(db *sqlx.DB) error {
	query := fmt.Sprintf(`
				CREATE TABLE IF NOT EXISTS %s (
			  	id SERIAL NOT NULL,
			  	description text NOT NULL,
			  	is_done boolean NOT NULL DEFAULT false,
			  	is_favourite boolean NOT NULL DEFAULT false,
			  	PRIMARY KEY (id)
			)`, "todo")
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	var todos []todo.Todo
	query = fmt.Sprintf("SELECT * FROM %s", "todo")
	err = db.Select(&todos, query)
	if err != nil {
		return err
	}
	if len(todos) != 0 {
		return nil
	}

	query = fmt.Sprintf("INSERT INTO %s (description, is_done, is_favourite) values ($1, $2, $3) RETURNING id", "todo")
	db.QueryRow(query, "todo1", false, false)
	db.QueryRow(query, "todo2", false, true)
	db.QueryRow(query, "todo3", true, false)
	db.QueryRow(query, "todo4", true, true)
	return nil
}
