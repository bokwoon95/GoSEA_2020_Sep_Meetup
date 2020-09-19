package ex4

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

func VanillaExample(db *sql.DB) {
	// simulate user input
	searchNameString := `%es%`
	searchEmailString := `%co%@email.com`
	timeStart := time.Time{}.Add(3 * time.Hour)
	timeEnd := time.Now()

	query := `SELECT u.user_id, u.name, u.email, u.created_at FROM public.users AS u WHERE `

	var predicates []string
	var args []interface{}

	// name
	if searchNameString != "" {
		predicates = append(predicates, fmt.Sprintf("u.name ILIKE ?"))
		args = append(args, searchNameString)
	}

	// email
	if searchEmailString != "" {
		predicates = append(predicates, fmt.Sprintf("u.email ILIKE ?"))
		args = append(args, searchEmailString)
	}

	// created_at
	if !timeStart.IsZero() && !timeEnd.IsZero() {
		predicates = append(predicates, fmt.Sprintf("u.created_at BETWEEN ? AND ?"))
		args = append(args, timeStart, timeEnd)
	}

	query = query + strings.Join(predicates, " AND ")

	// Rebind ?, ?, ? -> $1, $2, $3
	var newQuery string
	var count int
	for i := strings.Index(query, "?"); i >= 0; i = strings.Index(query, "?") {
		count++
		newQuery += query[:i]
		newQuery += fmt.Sprintf("$%d", count)
		query = query[i+1:]
	}

	// Print the results
	fmt.Println("query:", newQuery)
	fmt.Println("args:", args)
}
