package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/dpecos/cmdbox/models"
	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	db *sql.DB
)

func Init(dbPath string) {
	if dbPath == "" {
		log.Fatal("ERROR: cmdbox database path not specified")
	}
	os.Remove(dbPath)
	Load(dbPath)
	defer db.Close()

	log.Printf("cmdbox database successfully initialized in path %s\n", dbPath)
}

func Load(dbPath string) *sql.DB {
	if dbPath == "" {
		log.Fatal("ERROR: cmdbox database path not specified")
	}

	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	execSQL("PRAGMA foreign_keys = ON")

	updateSchema(dbPath)

	checkDefaultSpace(dbPath)

	return db
}

func updateSchema(dbPath string) {
	n, err := migrate.Exec(db, "sqlite3", migrations(), migrate.Up)
	if err != nil {
		log.Fatal(err)
	}
	if n != 0 {
		log.Printf("Applied %d migrations\n", n)
	}
}

func checkDefaultSpace(dbPath string) {
	spaces := SpacesList()
	if len(spaces) == 0 {
		id := SpacesCreate(models.Space{
			Name:  "default",
			Title: "Default space for commands",
		})
		log.Printf("Default space created %s\n", id)
	}
}

func execSQL(sql string, args ...interface{}) sql.Result {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
