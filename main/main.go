package main

import (
	"log"
	"net/http"

	"github.com/Rioh1118/cyoa"
)

func main() {
	jsonPath := "../gopher.json"
	templatePath := "../templates/format.gohtml"

	story, err := cyoa.ParseJson(jsonPath)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range story {
		outFilePath := "../templates/" + k + ".html"
		cyoa.MakeHtml(templatePath, outFilePath, v)
	}

	handleFunc := cyoa.MapHandler(story, "../templates")
	http.ListenAndServe(":8080", handleFunc)
}
