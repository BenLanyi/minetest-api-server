package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/benlanyi/go-dbtest/person"
)

type database struct {
	database *sqlx.DB
}

type returnMessage struct {
	Status  string
	Message string
}

// New : Constructor
func New() database {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "mysecretpassword"
		dbname   = "testdb"
	)

	connectionDetails := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println("Start")
	db, err := sqlx.Connect("postgres", connectionDetails)
	if err != nil {
		panic(err)
	}

	returnValue := database{db}
	return returnValue

}

// GetPerson : searches for a person by name and returns person object
func (db database) GetPerson(name string) []person.Person {
	people := []person.Person{}
	input := fmt.Sprintf("SELECT * FROM person WHERE name = '%s'", name)
	err := db.database.Select(&people, input)
	if err != nil {
		fmt.Print(err.Error())
	}
	return people
}

// AddPerson : adds a new record to the database with name and age
func (db database) AddPerson(name string, age int) returnMessage {

	people := []person.Person{}
	input := fmt.Sprintf("SELECT * FROM person WHERE name = '%s'", name)
	db.database.Select(&people, input)
	var returnMessage returnMessage
	if len(people) != 0 {
		fmt.Println("Record already exists: ", people)
		returnMessage.Status = "fail"
		returnMessage.Message = fmt.Sprintf("Record already exists with name %s", name)
	} else {
		input := fmt.Sprintf("INSERT INTO person (name, age) VALUES ('%s', %d)", name, age)
		db.database.MustExec(input)
		fmt.Println("Record inserted")
		returnMessage.Status = "success"
		returnMessage.Message = "Record inserted"
	}
	return returnMessage
}

// DeletePerson : deletes a record based off name and returns success/fail message
func (db database) DeletePerson(name string) returnMessage {
	people := []person.Person{}
	input := fmt.Sprintf("SELECT * FROM person WHERE name = '%s'", name)
	db.database.Select(&people, input)
	var returnMessage returnMessage
	if len(people) > 0 {
		input := fmt.Sprintf("DELETE FROM person WHERE name='%s'", name)
		db.database.MustExec(input)
		returnMessage.Status = "success"
		returnMessage.Message = "Record deleted"
	} else {
		returnMessage.Status = "fail"
		returnMessage.Message = fmt.Sprintf("Record with name %s not found. Can't delete", name)
	}
	return returnMessage
}

// GetAll : returns all records as person objects
func (db database) GetAll() []person.Person {
	people := []person.Person{}
	input := fmt.Sprintf("SELECT * FROM person")
	db.database.Select(&people, input)
	return people
}

// EditPerson : modifies the name or age of a person
func (db database) EditPerson(id int, name string, age int) returnMessage {
	input := fmt.Sprintf("UPDATE person SET name = '%s', age = %d WHERE ID = %d;", name, age, id)
	db.database.MustExec(input)
	fmt.Println("Record updated")
	var returnMessage returnMessage
	returnMessage.Status = "success"
	returnMessage.Message = "Record updated"
	return returnMessage
}
