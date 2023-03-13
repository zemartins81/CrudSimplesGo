package controller

import (
	"CrudSimplesGo/databases"
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("tmpl/*"))

func ListAll(w http.ResponseWriter, r *http.Request) {

	users := databases.SelectAll()
	err := tmpl.ExecuteTemplate(w, "Index", users)
	if err != nil {
		panic(err.Error())
	}
}

func ShowOne(w http.ResponseWriter, r *http.Request) {
	requestID := r.URL.Query().Get("id")

	user := databases.SelectOne(requestID)

	err := tmpl.ExecuteTemplate(w, "Show", user)
	if err != nil {
		panic(err.Error())
	}
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	email := r.FormValue("email")

	result := databases.InsertOne(name, email)
	log.Println(result)

	http.Redirect(w, r, "/", 301)

}
