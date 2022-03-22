package main

import (
	"context"
	"database/sql"
	"log"
	"reflect"

	_ "github.com/lib/pq"
	"github.com/rawansuww/prototypes/sqlc/tutorial"
)

func run() error {
	ctx := context.Background()

	db, err := sql.Open("postgres", "user=postgres dbname=sqlcdb sslmode=disable")
	if err != nil {
		return err
	}

	queries := tutorial.New(db)

	// list all Persons
	Persons, err := queries.ListPersons(ctx)
	if err != nil {
		return err
	}
	log.Println(Persons)

	// create an Person
	insertedPerson, err := queries.CreatePerson(ctx, tutorial.CreatePersonParams{
		Firstname: sql.NullString{String: "Brian", Valid: true},
		Lastname:  sql.NullString{String: "Kheeston", Valid: true},
		Email:     sql.NullString{String: "brian@golang.com", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertedPerson)

	// get the Person we just inserted
	fetchedPerson, err := queries.GetPerson(ctx, insertedPerson.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedPerson, fetchedPerson))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
