package groupie

import (
	"net/http"
	"text/template"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	tmpl.Execute(w, Struct)
}
