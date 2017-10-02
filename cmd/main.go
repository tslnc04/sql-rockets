package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "tslnc04"
	password = ""
	dbname   = "rockets"
)

func errHandle(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	pass, err := ioutil.ReadFile(".pg_auth")

	if err != nil {
		panic(err)
	}

	password = string(pass)
}

func main() {
	info := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", info)
	errHandle(err)
	defer db.Close()

	err = db.Ping()
	errHandle(err)

	rows, err := db.Query("SELECT name FROM engines WHERE thrust_vac > 100")
	errHandle(err)
	defer rows.Close()

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		errHandle(err)
		fmt.Println(name)
	}

	err = rows.Err()
	errHandle(err)

	fmt.Println("Success! Yay!")
}
