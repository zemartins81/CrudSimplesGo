package controller

import (
	"CrudSimplesGo/databases"
	"CrudSimplesGo/entities"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("tmpl/*"))

func Index(w http.ResponseWriter, r *http.Request) {

	db := databases.DBConn()
	defer db.Close()

	selDB, err := db.Query("SELECT * FROM names ORDER BY id DESC ")
	if err != nil {
		panic(err.Error())
	}

	n := entities.Names{}

	res := []entities.Names{}

	for selDB.Next() {
		var id int
		var name, email string

		err = selDB.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		n.Id = id
		n.Name = name
		n.Email = email

		res = append(res, n)
	}

	tmpl.ExecuteTemplate(w, "Index", res)
}
