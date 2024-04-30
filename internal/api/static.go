package api

import (
	"io/fs"
	"net/http"
	"text/template"

	"github.com/Brix101/nestfile/internal/domain"
	"go.uber.org/zap"
)

func (a api) indexHandler(w http.ResponseWriter, r *http.Request) {
	data := domain.StaticData{
		Name: "YourNameHere", // Replace "YourNameHere" with the desired name
	}

	fileContents, err := fs.ReadFile(a.assetsFs, "index.html")
	if err != nil {
		errorMessage := "failed to read file 'index.html' from assets filesystem"
		a.logger.Error(errorMessage, zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	index := template.Must(template.New("index").Delims("[{[", "]}]").Parse(string(fileContents)))
	err = index.Execute(w, data)
	if err != nil {
		errorMessage := "failed to execute template"
		a.logger.Error(errorMessage, zap.Error(err))
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return
	}
}
