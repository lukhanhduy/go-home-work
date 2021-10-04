package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"

	_ "github.com/lib/pq"
	pqsql "github.com/manabie-com/togo/internal/storages/pqsql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Print("on init")
	connStr := "user=compose-postgres dbname=compose-postgres sslmode=verify-full"

	fmt.Print("on init 1")
	db, err := sql.Open("postgres", connStr)

	fmt.Print("on init 2")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &pqsql.LiteDB{
			DB: db,
		},
	})
}
