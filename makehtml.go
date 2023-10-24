package cyoa

import (
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/format.gohtml"))
}

func MakeFile(outPath string, section StoryArc) error {

}
