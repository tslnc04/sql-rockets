package main

import (
	"fmt"

	"github.com/tslnc04/sql-rockets"

	_ "github.com/lib/pq"
)

var cfg *rockets.Config

func init() {
	cfg = rockets.LoadConfigFromFile(rockets.LoadFile(".pg_auth"))
}

func main() {
	info := rockets.NewConn(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBname)
	db := info.Connect()
	defer db.Close()

	rockets.TestPing(db)

	// interface of name
	rockets.AddRocket(db, "Test Name", 1.0, 1.0, "Test Manufacturer")
	iname := rockets.QueryDBRows(db, "SELECT name FROM rockets")

	// This code isn't exactly useful in this case, but could be
	var name []string
	for _, entry := range iname {
		name = append(name, entry.(string))
	}

	fmt.Println("Success! Yay!", name)
}
