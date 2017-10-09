package rockets

import "database/sql"

// AddRocket is an interactive tool to add a rocket to the table
func AddRocket(db *sql.DB, name string, height, diameter float32, manufacturer string) bool {
	statement := `INSERT INTO rockets (name, height, diameter, manufacturer)
VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(statement, name, height, diameter, manufacturer)

	if err != nil {
		panic(err)
	}

	return true
}

/*
Commented since rocket names are not necessarily unique
// FindRocketByName searches the rocket DB based on the name of the rocket
func FindRocketByName(db *sql.DB, name string) *Rocket {
	var output = new(Rocket)
	row := db.QueryRow("SELECT * FROM rockets WHERE name = $1;", name)

	err := row.Scan(&output.ID, &output.Name, &output.Height,
		&output.Diameter, &output.Manufacturer)

	if err != nil {
		panic(err)
	}

	return output
}
*/

// FindRocketByID searches the rocket DB based on the ID of the rocket
func FindRocketByID(db *sql.DB, id int) *Rocket {
	var output = new(Rocket)
	row := db.QueryRow("SELECT * FROM rockets WHERE id = $1", id)

	err := row.Scan(&output.ID, &output.Name, &output.Height,
		&output.Diameter, &output.Manufacturer)

	if err != nil {
		panic(err)
	}

	return output
}

func ChangeRocketManufacturer(db *sql.DB, id int, manu string) bool {
	statement := `UPDATE rockets SET manufacturer = $1 WHERE id = $2`
	_, err := db.Exec(statement, manu, id)

	if err != nil {
		panic(err)
	}

	return true
}
