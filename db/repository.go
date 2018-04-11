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

	return db
}

func Add(cmd models.Cmd) int64 {
	sqlStmt := `
	insert into commands(
		cmd, title, description, url
	) values ($1, $2, $3, $4)
	`
	result, err := db.Exec(sqlStmt, cmd.Cmd, cmd.Title, cmd.Description, cmd.URL)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func List() []models.Cmd {
	sqlStmt := `select * from commands`

	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	cmds := []models.Cmd{}
	for rows.Next() {
		var item models.Cmd
		if err := rows.Scan(&item.ID, &item.Cmd, &item.Title, &item.Description, &item.URL, &item.UpdatedAt, &item.CreatedAt); err != nil {
			log.Fatal(err)
		}

		item.Tags = Tags(item.ID)

		cmds = append(cmds, item)
	}

	return cmds
}

func Find(id int64) models.Cmd {
	sqlStmt := `select * from commands where id = $1`

	rows, err := db.Query(sqlStmt, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var item models.Cmd
	for rows.Next() {
		if err := rows.Scan(&item.ID, &item.Cmd, &item.Title, &item.Description, &item.URL, &item.UpdatedAt, &item.CreatedAt); err != nil {
			log.Fatal(err)
		}

		item.Tags = Tags(item.ID)
	}

	return item
}

func Tags(cmdID int64) []string {
	sqlStmt := `select tag from command_tags where command = $1`

	rows, err := db.Query(sqlStmt, cmdID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tags []string
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			log.Fatal(err)
		}

		tags = append(tags, tag)
	}

	return tags
}

func AssignTag(cmdID int64, tag string) {
	sqlStmt := `insert or ignore into tags(name) values ($1)`

	_, err := db.Exec(sqlStmt, tag)
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt = `insert or ignore into command_tags(command, tag) values ($1, $2)`

	_, err = db.Exec(sqlStmt, cmdID, tag)
	if err != nil {
		log.Fatal(err)
	}
}

func UnassignTag(cmdID int64, tag string) {
	execSQL(`delete from command_tags where command = $1 and tag = $2`)
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

func execSQL(sql string, args ...interface{}) {
	_, err := db.Exec(sql, args...)
	if err != nil {
		log.Fatal(err)
	}
}
