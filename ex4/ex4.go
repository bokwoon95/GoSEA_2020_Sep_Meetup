package ex4

import (
	"encoding/json"
	"fmt"
	"log"
)

func prettyprint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))
}
