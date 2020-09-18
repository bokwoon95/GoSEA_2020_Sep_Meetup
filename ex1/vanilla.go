package ex1

import (
	"database/sql"
	"log"
	"time"
)

func VanillaExample(db *sql.DB) {
	var users []User
	var user User

	rows, err := db.Query(`
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
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatalln(err)
	}

	// Print the results
	prettyprint(users)
}
