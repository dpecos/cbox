package db

import (
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
		var item models.Space
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			log.Fatal(err)
		}

		spaces = append(spaces, item)
	}

	return spaces
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
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			log.Fatal(err)
		}
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
		id, name
	) values ($1, $2)
	`
	execSQL(sqlStmt, id, space.Name)

	return id
}
