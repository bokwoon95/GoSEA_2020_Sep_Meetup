package ex2

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/bokwoon95/GoSEA_2020_Sep_Meetup/tables"
	sq "github.com/bokwoon95/go-structured-query/postgres"
)

type User struct {
	UserID    int
	Name      string
	Email     string
	CreatedAt time.Time
}

func (user *User) RowMapper(u tables.TABLE_USERS) func(*sq.Row) {
	return func(row *sq.Row) {
		user.UserID = row.Int(u.USER_ID)
		user.Name = row.String(u.NAME)
		user.Email = row.String(u.EMAIL)
		user.CreatedAt = row.Time(u.CREATED_AT)
	}
}

type SqlxUser struct {
	UserID    int       `db:"user_id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
}

func prettyprint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))
}
