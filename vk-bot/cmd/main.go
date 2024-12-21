package main

func main() {
	db := openConnection()
	longpollInit()
	closeConnection(db)
}
