package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/bokwoon95/GoSEA_2020_Sep_Meetup/ex1"
	"github.com/bokwoon95/GoSEA_2020_Sep_Meetup/ex2"
	"github.com/bokwoon95/GoSEA_2020_Sep_Meetup/ex3"
	"github.com/bokwoon95/GoSEA_2020_Sep_Meetup/ex3_modified"
	"github.com/bokwoon95/GoSEA_2020_Sep_Meetup/ex4"
	"github.com/bokwoon95/GoSEA_2020_Sep_Meetup/ex5"
	_ "github.com/lib/pq"
)

func main() {
	example := flag.String("example", "", "which example to run")
	flag.Parse()
	db, err := sql.Open("postgres", "postgres://pg:pg@localhost:5431/gosea_2020_sep_meetup?sslmode=disable")
	if err != nil {
		log.Fatalf("unable to open database connection: %s", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("unable to ping database: %s", err)
	}
	switch *example {
	// ex1
	case "1_vanilla":
		ex1.VanillaExample(db)
	case "1_sqlx":
		ex1.SqlxExample(db)
	case "1_sq":
		ex1.SqExample(db)

	// ex2
	case "2_vanilla":
		ex2.VanillaExample1(db)
		ex2.VanillaExample2(db)
		ex2.VanillaExample3(db)
	case "2_sqlx":
		ex2.SqlxExample1(db)
		ex2.SqlxExample2(db)
		ex2.SqlxExample3(db)
	case "2_sq":
		ex2.SqExample1(db)
		ex2.SqExample2(db)
		ex2.SqExample3(db)

	// ex3
	case "3_vanilla":
		ex3.VanillaExample(db)
	case "3_sq":
		ex3.SqExample(db)

	// ex3_modified
	case "3m_vanilla":
		ex3_modified.VanillaExample(db)
	case "3m_sqlx":
		ex3_modified.SqlxExample(db)
	case "3m_sq":
		ex3_modified.SqExample(db)

	// ex4
	case "4_vanilla":
		ex4.VanillaExample(db)
	case "4_sq":
		ex4.SqExample(db)

	// ex4
	case "5_vanilla":
		ex5.VanillaExample(db)
	case "5_sq":
		ex5.SqExample(db)
	}
}
