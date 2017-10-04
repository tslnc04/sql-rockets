package rockets

import "database/sql"

// AddRocket is an interactive tool to add a rocket to the table
func AddRocket(db *sql.DB, name string, height, diameter float32, manufacturer string) {
	statement := `INSERT INTO rockets (name, height, diameter, manufacturer)
VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(statement, name, height, diameter, manufacturer)

	if err != nil {
		panic(err)
	}
}
