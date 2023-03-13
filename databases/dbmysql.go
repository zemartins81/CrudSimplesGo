package databases

import (
	"CrudSimplesGo/entities"
	"database/sql"
)

func DBConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:33060)/database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}

func SelectAll() []entities.User {

	db := DBConn()
	defer db.Close()

	selDB, err := db.Query("SELECT * FROM users ORDER BY id DESC ")
	if err != nil {
		panic(err.Error())
	}

	n := entities.User{}

	var res []entities.User

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

	return res
}

func SelectOne(id string) entities.User {
	db := DBConn()
	defer db.Close()

	selected, err := db.Query("SELECT * FROM users WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}

	userEncontrado := entities.User{}

	var idEncontrado int
	var name, email string

	err = selected.Scan(&idEncontrado, &name, &email)
	if err != nil {
		panic(err.Error())
	}

	userEncontrado.Id = idEncontrado
	userEncontrado.Name = name
	userEncontrado.Email = email

	return userEncontrado
}

func InsertOne(name, email string) sql.Result {
	var addName = name
	var addEmail = email
	db := DBConn()
	defer db.Close()
	insUser, err := db.Prepare("INSERT INTO users(name, email) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}

	exec, err := insUser.Exec(addName, addEmail)
	if err != nil {
		panic(err.Error())
	}
	return exec

}
