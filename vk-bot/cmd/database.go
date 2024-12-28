package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

const (
	initStr1 = `
	CREATE TABLE IF NOT EXISTS users(
		id 			SERIAL PRIMARY KEY,
		FIO			TEXT,
		role 		TEXT 	DEFAULT user,
		chatStatus  INT 	DEFAULT 0,
		signed 		INT	 	DEFAULT 1
	);`
	initStr2 = `CREATE TABLE IF NOT EXISTS dialog(
		chatStatus  INT PRIMARY KEY,
		description	TEXT,
		payload		TEXT,
		keyboard	TEXT
	);`
	initStr3 = `CREATE TABLE IF NOT EXISTS events(
		id 				INT PRIMARY KEY,
		name 			TEXT,
		description 	TEXT,
		participants	INT
	);`
	newUserString = `
	INSERT INTO users 
	VALUES  ($1)
	`
	FIOString = `
	UPDATE users
	SET FIO = $2
	WHERE id = $1
	`
	changeRoleString = `
	UPDATE users
	SET role = $2
	WHERE id = $1
	`
	chatStatusUserString = `
	UPDATE users
	SET chatStatus = $2
	WHERE id = $1
	`
	getChatStatusString = `
	SELECT chatStatus
	FROM users
	WHERE id = $1
	LIMIT 1
	`
	getChatResponseString = `
	SELECT *
	FROM dialog
	WHERE chatStatus = $1
	LIMIT 1
	`
)

type ChatResponse struct {
	Text   string
	keyMap map[int]string
}

func openConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.host, cfg.port, cfg.user, cfg.password, cfg.dbname)
	var err error
	rnt.db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = rnt.db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to db is succesful")

	_, err = rnt.db.Exec(initStr1)
	if err != nil {
		log.Fatal(err)
	}
	_, err = rnt.db.Exec(initStr2)
	if err != nil {
		log.Fatal(err)
	}
	_, err = rnt.db.Exec(initStr3)
	if err != nil {
		log.Fatal(err)
	}
}

func newUser(id int) {
	_, err := rnt.db.Exec(newUserString, id)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("New user: %d\n", id)
	}
}

func FIOUser(id int, FIO string) {
	_, err := rnt.db.Exec(FIOString, id, FIO)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("FIO changed for %d to: %s\n", id, FIO)
	}
}

func changeRole(id int, role string) {
	_, err := rnt.db.Exec(changeRoleString, id, role)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("role changed for %d to: %s\n", id, role)
	}
}

func getChatResponse(chatStatus int) ChatResponse {
	var response ChatResponse
	response.keyMap = make(map[int]string)
	var keybuf string
	var payloadbuf string
	row := rnt.db.QueryRow(getChatResponseString, chatStatus)
	err := row.Scan(&chatStatus, &response.Text, &keybuf, &payloadbuf)
	if err != nil {
		log.Fatal(err)
	}
	keys := strings.Split(keybuf, ",")
	payload := strings.Split(payloadbuf, ",")
	for i, key := range keys {
		buf, err := strconv.Atoi(key)
		if err != nil {
			log.Fatal(err)
		}
		response.keyMap[buf] = payload[i]
	}
	log.Printf("Got response info for chatStatus = %d\ntext = %s keystring = %s payload = %s", chatStatus, response.Text, keybuf, payloadbuf)
	return response
}
