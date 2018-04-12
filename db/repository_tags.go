package db

import "log"

func TagsList() []string {
	sqlStmt := `select * from tags`

	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	tags := []string{}
	for rows.Next() {
		var item string
		if err := rows.Scan(&item); err != nil {
			log.Fatal(err)
		}

		tags = append(tags, item)
	}

	return tags
}

func TagsDelete(tag string) {
	sqlStmt := "delete from command_tags where tag = $1"
	execSQL(sqlStmt, tag)

	sqlStmt = "delete from tags where name = $1"
	execSQL(sqlStmt, tag)
}

func AssignTag(cmdID int64, tag string) {
	sqlStmt := `insert or ignore into tags(name) values ($1)`
	execSQL(sqlStmt, tag)

	sqlStmt = `insert or ignore into command_tags(command, tag) values ($1, $2)`
	execSQL(sqlStmt, cmdID, tag)
}

func UnassignTag(cmdID int64, tag string) {
	sqlStmt := `delete from command_tags where command = $1 and tag = $2`
	execSQL(sqlStmt, cmdID, tag)
}
