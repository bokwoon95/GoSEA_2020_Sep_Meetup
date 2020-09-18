package ex2

import (
	"database/sql"
	"log"
)

// Scenario 1: Get all users whose names start with "A"
func VanillaExample1(db *sql.DB) {
	var users []User
	var user User

	rows, err := db.Query(`
SELECT user_id, name, email, created_at
FROM users
WHERE SUBSTRING(name, 1, 1) = $1`,
		"A",
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
	prettyprint("Scenario 1: Get all users whose names start with 'A'")
	prettyprint(users)
}

// Scenario 2: Get all users whose user_ids are in the first 10 even numbers
func VanillaExample2(db *sql.DB) {
	var users []User
	var user User

	rows, err := db.Query(`
SELECT user_id, name, email, created_at
FROM users
WHERE user_id IN ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
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
	prettyprint("Scenario 2: Get all users whose user_ids are in the first 10 even numbers")
	prettyprint(users)
}

// Scenario 3: Get all users whose (name, email) tuple are in the list
func VanillaExample3(db *sql.DB) {
	var users []User
	var user User

	rows, err := db.Query(`
SELECT user_id, name, email, created_at
FROM users
WHERE (name, email) IN (($1, $2), ($3, $4), ($5, $6), ($7, $8), ($9, $10))`,
		"Tyler Bloggs", "tyler_bloggs@email.com",
		"Teddy Hayden", "teddy_hayden@email.com",
		"Ivy-Rose Mcdermott", "ivy-rose_mcdermott@email.com",
		"Danielle Gilmore", "danielle_gilmore@email.com",
		"Ronnie Kirby", "ronnie_kirby@email.com",
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
	prettyprint("Scenario 3: Get all users whose (name, email) tuple are in the list")
	prettyprint(users)
}
