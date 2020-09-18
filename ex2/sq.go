package ex2

import (
	"database/sql"
	"log"

	"github.com/bokwoon95/GoSEA_2020_Sep_Meetup/tables"
	sq "github.com/bokwoon95/go-structured-query/postgres"
)

// Scenario 1: Get all users whose names start with "A"
func SqExample1(db *sql.DB) {
	var users []User
	var user User

	u := tables.USERS()
	err := sq.
		WithDefaultLog(sq.Lverbose).
		Selectx(user.RowMapper(u), func() { users = append(users, user) }).
		From(u).
		Where(sq.Predicatef("SUBSTRING(?, 1, 1) = ?", u.NAME, "A")).
		Fetch(db)
	if err != nil {
		log.Fatalln(err)
	}

	// Print the results
	prettyprint("Scenario 1: Get all users whose names start with 'A'")
	prettyprint(users)
}

// Scenario 2: Get all users whose user_ids are in the first 10 even numbers
func SqExample2(db *sql.DB) {
	var users []User
	var user User

	u := tables.USERS()
	evenNumbers := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	err := sq.
		WithDefaultLog(sq.Lverbose).
		Selectx(user.RowMapper(u), func() { users = append(users, user) }).
		From(u).
		Where(u.USER_ID.In(evenNumbers)).
		Fetch(db)
	if err != nil {
		log.Fatalln(err)
	}

	// Print the results
	prettyprint("Scenario 2: Get all users whose user_ids are in the first 10 even numbers")
	prettyprint(users)
}

// Scenario 3: Get all users whose (name, email) tuple are in the list
func SqExample3(db *sql.DB) {
	var users []User
	var user User

	u := tables.USERS()
	tuples := sq.RowValues{
		sq.RowValue{"Tyler Bloggs", "tyler_bloggs@email.com"},
		sq.RowValue{"Teddy Hayden", "teddy_hayden@email.com"},
		sq.RowValue{"Ivy-Rose Mcdermott", "ivy-rose_mcdermott@email.com"},
		sq.RowValue{"Danielle Gilmore", "danielle_gilmore@email.com"},
		sq.RowValue{"Ronnie Kirby", "ronnie_kirby@email.com"},
	}
	err := sq.
		WithDefaultLog(sq.Lverbose).
		Selectx(user.RowMapper(u), func() { users = append(users, user) }).
		From(u).
		Where(sq.RowValue{u.NAME, u.EMAIL}.In(tuples)).
		Fetch(db)
	if err != nil {
		log.Fatalln(err)
	}

	// Print the results
	prettyprint("Scenario 3: Get all users whose (name, email) tuple are in the list")
	prettyprint(users)
}
