package ex1

import (
	"database/sql"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

func SqlxExample(db *sql.DB) {
	var dbx = sqlx.NewDb(db, "postgres")
	var users []SqlxUser

	err := dbx.Select(&users, `
SELECT
	user_id,
	name,
	email,
	created_at
FROM users
WHERE
	name <> email
	AND created_at < $1
	AND deleted_at IS NULL
LIMIT $2`,
		time.Now(),
		200,
	)
	if err != nil {
		log.Fatalln(err)
	}

	// Print the results
	prettyprint(users)
}
