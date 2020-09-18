package ex3_modified

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

func SqlxExample(db *sql.DB) {
	var dbx = sqlx.NewDb(db, "postgres")
	var submissions []SqlxSubmission

	err := dbx.Select(&submissions, `
SELECT
	-- submissions
	s.submission_id,
	s.answers,
	s.submitted,

	-- assignments
	a.assignment_id,
	a.questions,

	-- teams
	t.team_id,
	t.team_name,

	-- user
	u.user_id,
	u.name,
	u.email,
	u.created_at
FROM
	submissions AS s
	JOIN assignments AS a ON a.assignment_id = s.assignment_id
	JOIN teams AS t ON t.team_id = s.team_id
	JOIN users AS u ON u.team_id = t.team_id`,
	)
	if err != nil {
		log.Fatalln(err)
	}

	// Print the results
	// prettyprint(submissions)
}
