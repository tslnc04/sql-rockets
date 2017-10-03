package main

import (
	"fmt"
	"io/ioutil"

	"github.com/tslnc04/sql-rockets"

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
	info := rockets.NewConn(host, port, user, password, dbname)
	db := info.Connect()
	defer db.Close()

	rockets.TestPing(db)

	rows, err := db.Query("SELECT * FROM engines WHERE thrust_vac > 100")
	errHandle(err)
	defer rows.Close()

	for rows.Next() {
		var name = make([]interface{}, 9)
		err = rows.Scan(&name[0], &name[1], &name[2], &name[3], &name[4], &name[5], &name[6], &name[7], &name[8])
		errHandle(err)
		fmt.Println(name)
	}

	err = rows.Err()
	errHandle(err)

	fmt.Println("Success! Yay!")
}
