package rockets

import (
	"database/sql"
	"fmt"
)

type conn struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

// NewConn creates a new conn struct with the information given
func NewConn(host string, port int, user, password, dbname string) *conn {
	newConn := conn{host, port, user, password, dbname}

	return &newConn
}

func (c *conn) Connect() *sql.DB {
	info := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		c.host, c.port, c.user, c.password, c.dbname)

	db, err := sql.Open("postgres", info)

	if err != nil {
		panic(err)
	}

	return db
}

// TestPing pings the db and returns true if it was successful, panicking otherwise
func TestPing(db *sql.DB) bool {
	err := db.Ping()

	if err != nil {
		panic(err)
	}

	return true
}
