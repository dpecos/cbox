package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/dpecos/cmdbox/models"
	"github.com/mattes/migrate"
	sqlite3 "github.com/mattes/migrate/database/sqlite3"
	_ "github.com/mattes/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
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
	log.Printf("Using database: %s", dbPath)

	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	updateSchema(dbPath)

	return db
}

func Add(cmd models.Cmd) {
	sqlStmt := `
	insert into commands(
		cmd, title, description, url
	) values ($1, $2, $3, $4)
	`
	_, err := db.Exec(sqlStmt, cmd.Cmd, cmd.Title, cmd.Description, cmd.URL)
	if err != nil {
		log.Fatal(err)
	}
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

func Tags(cmdID int) []string {
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

func AssignTag(cmdID int, tag string) {
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

func UnassignTag(cmdID int, tag string) {
	sqlStmt := `delete from command_tags where command = $1 and tag = $2`

	_, err := db.Exec(sqlStmt, cmdID, tag)
	if err != nil {
		log.Fatal(err)
	}
}

func updateSchema(dbPath string) {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://./db/migrations", "ql", driver)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
