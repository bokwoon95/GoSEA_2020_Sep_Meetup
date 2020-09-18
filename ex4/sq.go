package ex4

import (
	"database/sql"
	"log"

	"github.com/bokwoon95/GoSEA_2020_Sep_Meetup/tables"
	sq "github.com/bokwoon95/go-structured-query/postgres"
)

func SqExample(db *sql.DB) {
	var submissions []Submission
	var submission Submission

	vs := tables.V_SUBMISSIONS()
	err := sq.
		WithDefaultLog(sq.Lverbose).
		Selectx(submission.RowMapper(vs), func() {
			submissions = append(submissions, submission)
		}).
		From(vs).
		Fetch(db)
	if err != nil {
		log.Fatalln(err)
	}

	// Print the results
	prettyprint(submissions)
}
