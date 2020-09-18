package ex3_modified

import (
	"database/sql"
	"log"

	"github.com/bokwoon95/GoSEA_2020_Sep_Meetup/tables"
	sq "github.com/bokwoon95/go-structured-query/postgres"
)

func SqExample(db *sql.DB) {
	var submissions []Submission
	var submission Submission

	s := tables.SUBMISSIONS().As("s")
	a := tables.ASSIGNMENTS().As("a")
	t := tables.TEAMS().As("t")
	u := tables.USERS().As("u")
	err := sq.
		// WithDefaultLog(sq.Lverbose).
		From(s).
		Join(a, a.ASSIGNMENT_ID.Eq(s.ASSIGNMENT_ID)).
		Join(t, t.TEAM_ID.Eq(s.TEAM_ID)).
		Join(u, u.TEAM_ID.Eq(t.TEAM_ID)).
		Selectx(
			submission.RowMapper(s, a, t, u),
			func() { submissions = append(submissions, submission) },
		).
		Fetch(db)
	if err != nil {
		log.Fatalln(err)
	}

	// Print the results
	// prettyprint(submissions)
}
