package main

import (
	"os"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type person struct {
	Id int
	LastName string
	FirstName string
	Address string
	City string
}

func main() {

	fmt.Println("Opening connection")
	db, err := sql.Open("mysql", "root:"+os.Getenv("mySQLPassword")+"@tcp("+os.Getenv("mySQLIPAddress")+":"+ os.Getenv("mySQLIPPort") +")/")

	if err != nil {
		fmt.Println("Error: ", err)
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	fmt.Println("Connection opened successfully!")

	// Last thing to do, close connection
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Execute the query
	var (
		Id int
		LastName string
		FirstName string
		Address string
		City string
	)

	//// Create Database if not exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS testdb")
	if err!=nil{
		log.Fatal(err)
	}

	//// Use our database
	_, err = db.Exec("USE testdb")
	if err!=nil{
		log.Fatal(err)
	}

	//// Create Table if not exists
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Persons (PersonID int NOT NULL AUTO_INCREMENT, LastName varchar(255), FirstName varchar(255), Address varchar(255), City varchar(255), PRIMARY KEY (PersonID));")
	if err!=nil{
		log.Fatal(err)
	}

	///////////// Read Operation ///////////////////
	rowsQuery, err := db.Prepare("SELECT * FROM Persons where PersonID = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rowsQuery.Close()
	// Parametrize
	rows, err := rowsQuery.Query(1)
	for rows.Next() {
		err := rows.Scan(&Id,  &LastName, &FirstName, &Address, &City)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(Id, LastName, FirstName, Address, City)
	}

	//// Get column names
	//columns, err := rows.Columns()
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	//fmt.Println(columns)

	//////////// Insertion ////////////////////////
	stmt, err := db.Prepare("INSERT INTO Persons(LastName, FirstName, Address, City) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	// Single Record Addition
	res, err := stmt.Exec("Khalid", "Butt", "ABC", "Lahore")
	if err != nil {
		log.Fatal(err)
	}

	// For MultiInsertion Case struct can be use that can be dumped with data from any source like: var persons = []person{}

	// META
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	///////////// Delete a Record ////////////////
	//_, err = db.Exec("DELETE FROM Persons where PersonID = ?", 2)  // OK
	//if err!=nil{
	//	log.Fatal(err)
	//}

	///////////// Update a Record ////////////////
	//_, err = db.Exec("UPDATE Persons SET City = 'Karachi' where PersonID = ?", 2)  // OK
	//if err!=nil{
	//	log.Fatal(err)
	//}

}
