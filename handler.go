package rockets

import (
	"database/sql"
)

// QueryDBRow uses a bit of sql to return queries
func QueryDBRow(db *sql.DB, query string) interface{} {
	var output interface{}
	row := db.QueryRow(query)

	err := row.Scan(&output)

	if err != nil {
		panic(err)
	}

	return output
}

// QueryDBRows queries and returns multiple rows of the database
func QueryDBRows(db *sql.DB, query string) []interface{} {
	output := []interface{}{}
	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var singleOutput interface{}
		err := rows.Scan(&singleOutput)
		output = append(output, singleOutput)

		if err != nil {
			panic(err)
		}
	}

	return output
}
