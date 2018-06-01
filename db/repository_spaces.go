package db

import (
	"database/sql"
	"log"

	"github.com/satori/go.uuid"

	"github.com/dpecos/cmdbox/models"
)

func SpacesList() []models.Space {
	sqlStmt := `select * from spaces`

	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	spaces := []models.Space{}
	for rows.Next() {
		item := extractItemFromRow(rows)
		spaces = append(spaces, item)
	}

	return spaces
}

func extractItemFromRow(rows *sql.Rows) models.Space {
	var item models.Space
	if err := rows.Scan(&item.ID, &item.Name, &item.Title); err != nil {
		log.Fatal(err)
	}
	return item
}

func SpacesFind(id uuid.UUID) models.Space {
	sqlStmt := `select * from spaces where id = $1`

	rows, err := db.Query(sqlStmt, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var item models.Space
	for rows.Next() {
		item = extractItemFromRow(rows)
	}

	if item.Name == "" {
		log.Fatalf("Space with ID %s not found", id)
	}

	return item
}

func SpacesCreate(space models.Space) uuid.UUID {
	id := uuid.Must(uuid.NewV4())

	sqlStmt := `
	insert into spaces(
		id, name, title
	) values ($1, $2, $3)
	`
	execSQL(sqlStmt, id, space.Name, space.Title)

	return id
}

func SpacesDelete(id uuid.UUID) {
	sqlStmt := "delete from spaces where id = $1"
	execSQL(sqlStmt, id)
}

func SpacesUpdate(space models.Space) {
	sqlStmt := `update spaces set 
		name = $1,
		title = $2
	where id = $3
	`
	execSQL(sqlStmt, space.Name, space.Title, space.ID)
}
