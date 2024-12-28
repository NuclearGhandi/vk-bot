package main

import (
	"database/sql"

	"github.com/SevereCloud/vksdk/v3/api"
)

type Config struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	VKTOKEN  string
}

type Runtime struct {
	db *sql.DB
	vk *api.VK
}

var cfg Config

var rnt Runtime

func main() {
	cfg.host = "localhost"
	cfg.port = 5432
	cfg.user = "postgres"
	cfg.password = "seliger"
	cfg.dbname = "vk-local"
	cfg.VKTOKEN = "vk1.a.8RlrGVwJCs7TSx55OaJHKvTTnGFCZ_fiba-t9rWJWIRtRDuxSWmH-Vlh18Np7jiRLur5-4suJNbThrT9k17RMZRjMMpe0qShwvfvv78xBIg9JCcQ2PG3NXjPegil18ZBWrex3EIL67xiDjb_ZNcbO4B-8B_IPDhR0d_kJVjWCAnH4f4P3MxZn2fq9x_wRobVZ8BRyoYhWsNPZQmzZFqmJw"
	openConnection()
	getChatResponse(1)
	defer rnt.db.Close()

	//longpollInit()
}
