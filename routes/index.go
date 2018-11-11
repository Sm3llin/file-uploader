package routes

import (
	"github.com/sm3llin/file-uploader/templates"
	"html/template"
	"net/http"
)

// Main route for the webapp, will serve the static html page to the user.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	str, err := templates.Box.FindString("index.html")
	if err != nil {
		panic(err)
	}

	t, err := template.ParseGlob(str)
	if err != nil {
		panic(t)
	}
	t.Execute(w, nil)
}
