package ex3

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
	u1, u2 := tables.USERS().As("u1"), tables.USERS().As("u2")
	err := sq.
		WithDefaultLog(sq.Lverbose).
		From(s).
		Join(a, a.ASSIGNMENT_ID.Eq(s.ASSIGNMENT_ID)).
		Join(t, t.TEAM_ID.Eq(s.TEAM_ID)).
		Join(u1, u1.TEAM_ID.Eq(t.TEAM_ID)).
		Join(u2, u2.TEAM_ID.Eq(t.TEAM_ID), u2.USER_ID.Gt(u1.USER_ID)).
		Selectx(submission.RowMapper(s, a, t, u1, u2), func() {
			submissions = append(submissions, submission)
		}).
		Limit(10).
		Fetch(db)
	if err != nil {
		log.Fatalln(err)
	}

	// Print the results
	prettyprint(submissions)
}
