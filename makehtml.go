package cyoa

import (
	"errors"
	"html/template"
	"os"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/format.gohtml"))
}

func MakeHtml(fileName string, key string, story Story) error {
	if !strings.HasSuffix(fileName, ".html") {
		return errors.New("file name isn't an html file")
	}
	path := "templates/" + fileName

	nf, err := os.Create(path)
	if err != nil {
		return err
	}
	defer nf.Close()

	err = tpl.ExecuteTemplate(nf, "templates/format.gohtml", story[key])
	if err != nil {
		return err
	}
	return nil
}
