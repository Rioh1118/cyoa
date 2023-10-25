package cyoa

import (
	"errors"
	"html/template"
	"os"
	"strings"
)

var tpl *template.Template

func MakeHtml(fileName string, key string, story Story) error {
	// 指定されたファイル名が.htmlで終わっているかを確認
	if !strings.HasSuffix(fileName, ".html") {
		return errors.New("file name isn't an html file")
	}

	// テンプレートのパスを指定
	templatePath := "templates/format.gohtml"

	// テンプレートをパース
	tpl = template.Must(template.ParseFiles(templatePath))

	// 出力ファイルを作成
	outputFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// テンプレートを実行し、結果を出力ファイルに書き込む
	err = tpl.ExecuteTemplate(outputFile, templatePath, story[key])
	if err != nil {
		return err
	}

	return nil
}
