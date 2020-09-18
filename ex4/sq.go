package ex4

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bokwoon95/GoSEA_2020_Sep_Meetup/tables"
	sq "github.com/bokwoon95/go-structured-query/postgres"
)

func SqExample(db *sql.DB) {
	names := []string{
		"Kaylie Stone",
		"Rhiann Hunter",
		"Teddy Hayden",
		"Lisa-Marie Mccormack",
		"Misbah Drake",
		"Dan Connelly",
		"Izaac Jeffery",
		"Lee Moon",
		"Lily Ray",
		"Jasmin Pope",
		"Arvin Kennedy",
		"Ayyub Cortez",
		"Krystian Boyce",
		"Zack Bell",
		"Lincoln Stevens",
		"Beck Coleman",
		"Sam Bevan",
		"Danielle Gilmore",
		"Salman Feeney",
	}

	emails := []string{
		"kaylie_stone@email.com",
		"rhiann_hunter@email.com",
		"teddy_hayden@email.com",
		"lisa-marie_mccormack@email.com",
		"misbah_drake@email.com",
		"dan_connelly@email.com",
		"izaac_jeffery@email.com",
		"lee_moon@email.com",
		"lily_ray@email.com",
		"jasmin_pope@email.com",
		"arvin_kennedy@email.com",
		"ayyub_cortez@email.com",
		"krystian_boyce@email.com",
		"zack_bell@email.com",
		"lincoln_stevens@email.com",
		"beck_coleman@email.com",
		"sam_bevan@email.com",
		"danielle_gilmore@email.com",
		"salman_feeney@email.com",
	}

	u := tables.USERS()
	q := sq.InsertInto(u).Columns(u.NAME, u.EMAIL).OnConflict().DoNothing()
	for i := range names {
		q = q.Values(names[i], emails[i])
	}
	_, err := q.Exec(db, 0)
	if err != nil {
		log.Fatalln(err)
	}

	// Print the results
	query, args := q.ToSQL()
	fmt.Println("query:", query)
	fmt.Println()
	fmt.Println("args:", args)
}
