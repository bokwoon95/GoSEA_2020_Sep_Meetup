package ex4

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func VanillaExample(db *sql.DB) {
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

	var placeholders []string
	var args []interface{}
	for i := range names {
		// i=0: ($1, $2)
		// i=1: ($3, $4)
		// i=2: ($5, $6)
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d)", (i*2)+1, (i+1)*2))
		args = append(args, names[i], emails[i])
	}
	query := "INSERT INTO public.users (name, email) VALUES " + strings.Join(placeholders, ", ") + " ON CONFLICT DO NOTHING"
	_, err := db.Exec(query, args...)
	if err != nil {
		log.Fatalln(err)
	}

	// Print the results
	fmt.Println("query:", query)
	fmt.Println()
	fmt.Println("args:", args)
}
