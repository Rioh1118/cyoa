package cyoa

import (
	"io"
	"net/http"
	"os"
	"strings"
)

/*
map
index intro.html => /
other key.html => /key
*/
func MapHandler(story Story, dirPath string) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			key := strings.TrimPrefix(path, "/")

			if _, ok := story[key]; !ok {
				http.NotFound(w, r)
				return
			}
			filePath := dirPath + "/" + path + ".html"

			file, err := os.Open(filePath)
			if err != nil {
				http.Error(w, "File not found", http.StatusNotFound)
				return
			}
			defer file.Close()

			_, err = io.Copy(w, file)
			if err != nil {
				http.Error(w, "Failed to send response", http.StatusInternalServerError)
				return
			}
		})
}
