package cyoa

import (
	"errors"
	"html/template"
	"os"
	"strings"
)

var tpl *template.Template

func MakeHtml(tempPath string, filePath string, chapter Chapter) error {
	// 指定されたファイル名が.htmlで終わっているかを確認
	if !strings.HasSuffix(filePath, ".html") {
		return errors.New("file name isn't an html file")
	}

	// テンプレートをパース
	tpl = template.Must(template.ParseFiles(tempPath))

	// 出力ファイルを作成
	outputFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// テンプレートを実行し、結果を出力ファイルに書き込む
	err = tpl.ExecuteTemplate(outputFile, tempPath, chapter)
	if err != nil {
		return err
	}

	return nil
}
