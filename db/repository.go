package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/dpecos/cmdbox/models"
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
	createSchema()
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
		cmds = append(cmds, item)
	}

	return cmds
}

func createSchema() {
	sqlStmt := `
	create table commands (
		id integer not null primary key,
		cmd text not null,
		title text not null,
		description text,
		url text,
		created_at timestamp default current_timestamp,
		updated_at timestamp default current_timestamp
	);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}
