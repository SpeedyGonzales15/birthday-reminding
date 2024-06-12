package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

var DB *sql.DB

func init() {
	const (
		host     = "localhost"
		port     = "5432"
		user     = "postgres"
		password = "postgres"
		dbname   = "task_db"
		sslmode  = "disable"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	goose.SetDialect("postgres")
	goose.SetBaseFS(embedMigrations)

	if err := goose.Up(DB, "migrations"); err != nil { //
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
func main() {

}
