package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
		return
	}

	var catid int
	err = db.QueryRow(`INSERT INTO cats(name) VALUES('beatrice') RETURNING id`).Scan(&catid)

	log.Printf("Created %d", catid)

	if err != nil {
		log.Fatal(err)
	}
}
