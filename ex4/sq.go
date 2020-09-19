package ex4

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/bokwoon95/GoSEA_2020_Sep_Meetup/tables"
	sq "github.com/bokwoon95/go-structured-query/postgres"
)

func SqExample(db *sql.DB) {
	// simulate user input
	searchNameString := `%es%`
	searchEmailString := `%co%@email.com`
	timeStart := time.Time{}.Add(3 * time.Hour)
	timeEnd := time.Now()

	u := tables.USERS().As("u")
	q := sq.Select(u.USER_ID, u.NAME, u.EMAIL, u.CREATED_AT).From(u)

	// name
	if searchNameString != "" {
		q = q.Where(u.NAME.ILikeString(searchNameString))
	}

	// email
	if searchEmailString != "" {
		q = q.Where(u.EMAIL.ILikeString(searchEmailString))
	}

	// created_at
	if !timeStart.IsZero() && !timeEnd.IsZero() {
		q = q.Where(u.CREATED_AT.BetweenTime(timeStart, timeEnd))
	}

	// Print the results
	query, args := q.ToSQL()
	fmt.Println("query:", query)
	fmt.Println("args:", args)
}
