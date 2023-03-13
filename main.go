package main

import (
	"CrudSimplesGo/controller"
	"net/http"
)

func main() {

	// Gerencia as URLs
	http.HandleFunc("/", controller.ListAll)
	http.HandleFunc("/show", controller.ShowOne)
	http.HandleFunc("/new", add)
	http.HandleFunc("/edit", Edit)

	// Ações
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

	// Inicia o servidor na porta 9000
	http.ListenAndServe(":9000", nil)
}
