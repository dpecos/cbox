package db

import (
	"log"

	"github.com/dpecos/cmdbox/models"
)

func Add(cmd models.Cmd) int64 {
	sqlStmt := `
	insert into commands(
		cmd, title, description, url
	) values ($1, $2, $3, $4)
	`
	result := execSQL(sqlStmt, cmd.Cmd, cmd.Title, cmd.Description, cmd.URL)

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func Delete(id int64) {
	sqlStmt := "delete from command_tags where command = $1"
	execSQL(sqlStmt, id)

	sqlStmt = "delete from commands where id = $1"
	execSQL(sqlStmt, id)
}

func List(tag string) []models.Cmd {
	sqlStmt := `select * from commands`
	if tag != "" {
		sqlStmt = `select c.* from commands c join command_tags ct on ct.command = c.ID where ct.tag = $1`
	}

	rows, err := db.Query(sqlStmt, tag)
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

		item.Tags = commandTags(item.ID)

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

		item.Tags = commandTags(item.ID)
	}

	if item.ID == 0 {
		log.Fatalf("Command with id %d not found", id)
	}

	return item
}

func Update(cmd models.Cmd) {
	sqlStmt := `update commands set 
		cmd = $1,
		title = $2,
		description = $3,
		URL = $4
	where id = $5
	`
	execSQL(sqlStmt, cmd.Cmd, cmd.Title, cmd.Description, cmd.URL, cmd.ID)
}

func commandTags(cmdID int64) []string {
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
