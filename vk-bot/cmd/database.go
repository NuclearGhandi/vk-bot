package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host       = "localhost"
	port       = 5432
	user       = "postgres"
	password   = "your-password"
	dbname     = "calhounio_demo"
	initString = `
	CREATE TABLE IF NOT EXISTS users(
		id 			SERIAL KEY,
		nametag		TEXT,
		role 		TEXT 	DEFAULT user,
		chatStatus  INT 	DEFAULT 0,
		signed 		INT	 	DEFAULT 1
	);
	CREATE TABLE IF NOT EXISTS dialog(
		chatStatus  INT KEY,
		description	TEXT,
		payload		TEXT,
		keyboard	TEXT,
	);
	CREATE TABLE IF NOT EXISTS events(
		id 				INT KEY,
		name 			TEXT,
		description 	TEXT,
		participants	INT
	);
	`
	newUserString = `
	INSERT INTO users 
	VALUES ($1, $2)
	`
	changeRoleString = `
	UPDATE users
	VALUES ($1)
	`
)

func openConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to db is succesful")

	_, err = db.Exec(initString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func newUser()

func closeConnection(db *sql.DB) {
	db.Close()
}
