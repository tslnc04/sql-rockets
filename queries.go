package rockets

import (
	"database/sql"
	"log"
)

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
	row := db.QueryRow("SELECT * FROM rockets WHERE id = $1;", id)

	err := row.Scan(&output) /*.ID, &output.Name, &output.Height,
	&output.Diameter, &output.Manufacturer)*/

	if err == sql.ErrNoRows {
		log.Fatal("No results found")
	}

	if err != nil {
		panic(err)
	}

	return output
}

// ChangeRocketManufacturer updates the database changing manufacturer
func ChangeRocketManufacturer(db *sql.DB, id int, manu string) bool {
	statement := `UPDATE rockets SET manufacturer = $1 WHERE id = $2;`
	_, err := db.Exec(statement, manu, id)

	if err != nil {
		panic(err)
	}

	return true
}

// AddOrUpdateRocket preforms an upsert adding or modifying the existing id
func AddOrUpdateRocket(db *sql.DB, id int, name string, height, diameter float32, manufacturer string) bool {
	statement := `INSERT INTO rockets (id, name, height, diameter, manufacturer)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (id) DO UPDATE SET id = $1, name = $2, height = $3, diameter = $4, manufacturer = $5;`
	_, err := db.Exec(statement, id, name, height, diameter, manufacturer)

	if err != nil {
		panic(err)
	}

	return true
}

// FindStageEngines returns 'limit' number of entries consisting of the
// stage's id, the engine's id, and the name of the engine
func FindStageEngines(db *sql.DB, limit int) ([]int, []int, []string) {
	statement := `
    SELECT stages.stage_id, stages.engine_id, engines.name
      FROM stages
RIGHT JOIN engines
        ON stages.engine_id = engines.id
     LIMIT $1;`

	rows, err := db.Query(statement, limit)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var outputStageID []int
	var outputEngineID []int
	var outputName []string

	for rows.Next() {
		var stageID sql.NullInt64
		var engineID sql.NullInt64
		var name string

		err = rows.Scan(&stageID, &engineID, &name)

		if err != nil {
			panic(err)
		}

		outputStageID = append(outputStageID, int(stageID.Int64))
		outputEngineID = append(outputEngineID, int(engineID.Int64))
		outputName = append(outputName, name)
	}

	err = rows.Err()

	if err != nil {
		panic(err)
	}

	return outputStageID, outputEngineID, outputName
}

// FindStageRockets returns 'limit' number of entries consisting of the
// stage's id, the rocket's id, and the name of the
func FindStageRockets(db *sql.DB, limit int) ([]int, []int, []string) {
	statement := `
    SELECT stages.stage_id, stages.rocket_id, rockets.name
      FROM stages
RIGHT JOIN rockets
        ON stages.rocket_id = rockets.id
     LIMIT $1;`

	rows, err := db.Query(statement, limit)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var outputStageID []int
	var outputRocketID []int
	var outputName []string

	for rows.Next() {
		var stageID sql.NullInt64
		var rocketID sql.NullInt64
		var name string

		err = rows.Scan(&stageID, &rocketID, &name)

		if err != nil {
			panic(err)
		}

		outputStageID = append(outputStageID, int(stageID.Int64))
		outputRocketID = append(outputRocketID, int(rocketID.Int64))
		outputName = append(outputName, name)
	}

	err = rows.Err()

	if err != nil {
		panic(err)
	}

	return outputStageID, outputRocketID, outputName
}
