package databases

import "database/sql"

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
