package main

import (
	"fmt"
	"log"

	"github.com/Rioh1118/cyoa"
)

func main() {
	story, err := cyoa.ParseJson("cyoa/templates/gopher.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(story)
}
