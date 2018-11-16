package service

import (
	"net/http"
	"text/template"

	"github.com/unrolled/render"
)

func homeHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "index", struct {
			ID      string `json:"id"`
			Content string `json:"content"`
		}{ID: "8675309", Content: "Hello from Go!"})
	}
}

func homepage(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	username := req.Form["username"][0]
	password := req.Form["password"][0]

	t := template.Must(template.ParseFiles("templates/table.html"))
	t.Execute(w, map[string]string{
		"username": username,
		"password": password,
	})
}

func NotImplemented(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "501 Not Implemented", 501)
}
