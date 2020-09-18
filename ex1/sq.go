package ex1

import (
	"database/sql"
	"log"
	"time"

	"github.com/bokwoon95/GoSEA_2020_Sep_Meetup/tables"
	sq "github.com/bokwoon95/go-structured-query/postgres"
)

func SqExample(db *sql.DB) {
	var users []User
	var user User

	u := tables.USERS()
	err := sq.
		WithDefaultLog(sq.Lverbose).
		Selectx(func(row *sq.Row) {
			user.UserID = row.Int(u.USER_ID)
			user.Name = row.String(u.NAME)
			user.Email = row.String(u.EMAIL)
			user.CreatedAt = row.Time(u.CREATED_AT)
		}, func() {
			users = append(users, user)
		}).
		From(u).
		Where(
			u.EMAIL.Ne(u.NAME),
			u.CREATED_AT.LtTime(time.Now()),
			u.DELETED_AT.IsNull(),
		).
		Limit(200).
		Fetch(db)
	if err != nil {
		log.Fatalln(err)
	}

	// Print the results
	prettyprint(users)
}
