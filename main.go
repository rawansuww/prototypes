package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

var schema = `
DROP TABLE IF EXISTS person;
CREATE TABLE person (
    first_name text,
    last_name text,
    email text
);`

func main() { //adjust user and db name!
	db, err := sqlx.Connect("postgres", "user=postgres dbname=sqlxdb sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	db.MustExec(schema)

	tx := db.MustBegin() //insert statements
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Noora", "Hussein", "noora@hotmail.com")
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Dianna", "Spencer", "queen@hotmail.com")
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Monicca", "Bellucci", "jesus@hotmail.com")
	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	tx.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Person{"Jane", "Citizen", "jane.citzen@example.com"})
	tx.Commit()

	// Query the database, storing results in a []Person (wrapped in []interface{})
	people := []Person{}
	db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
	jason, john := people[0], people[1]

	fmt.Printf("%#v\n%#v\n", jason, john)

	// if you have null fields and use SELECT *, you must use sql.Null* in your struct
	morePeople := []Person{}
	err = db.Select(&morePeople, "SELECT * FROM person ORDER BY email ASC")
	if err != nil {
		fmt.Println(err)
		return
	}
	peep1, peep2, peep3 := morePeople[0], morePeople[1], morePeople[2]

	fmt.Printf("%#v\n%#v\n%#v\n", peep1, peep2, peep3)
	// Loop through rows using only one struct
	pop := Person{}
	rows, err := db.Queryx("SELECT * FROM person")
	for rows.Next() {
		err := rows.StructScan(&pop)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", pop)
	}
	// Named queries, using `:name` as the bindvar.  Automatic bindvar support
	// which takes into account the dbtype based on the driverName on sqlx.Open/Connect
	_, err = db.NamedExec(`INSERT INTO person (first_name,last_name,email) VALUES (:first,:last,:email)`,
		map[string]interface{}{
			"first": "Bin",
			"last":  "Smuth",
			"email": "bensmith@allblacks.nz",
		})

	// Selects Mr. Smith from the database
	rows, err = db.NamedQuery(`SELECT * FROM person WHERE first_name=:fn`, map[string]interface{}{"fn": "Bin"})

	// Named queries can also use structs.  Their bind names follow the same rules
	// as the name -> db mapping, so struct fields are lowercased and the `db` tag
	// is taken into consideration.
	rows, err = db.NamedQuery(`SELECT * FROM person WHERE first_name=:first_name`, jason)

	// batch insert

	// batch insert with structs
	personStructs := []Person{
		{FirstName: "Ardie", LastName: "Savea", Email: "asavea@ab.co.nz"},
		{FirstName: "Sonny Bill", LastName: "Williams", Email: "sbw@ab.co.nz"},
		{FirstName: "Ngani", LastName: "Laumape", Email: "nlaumape@ab.co.nz"},
	}

	_, err = db.NamedExec(`INSERT INTO person (first_name, last_name, email)
        VALUES (:first_name, :last_name, :email)`, personStructs)

	// batch insert with maps
	personMaps := []map[string]interface{}{
		{"first_name": "Ardie", "last_name": "Savea", "email": "asavea@ab.co.nz"},
		{"first_name": "Sonny Bill", "last_name": "Williams", "email": "sbw@ab.co.nz"},
		{"first_name": "Ngani", "last_name": "Laumape", "email": "nlaumape@ab.co.nz"},
	}

	_, err = db.NamedExec(`INSERT INTO person (first_name, last_name, email)
        VALUES (:first_name, :last_name, :email)`, personMaps)
}
