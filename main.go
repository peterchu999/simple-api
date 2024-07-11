package main

import (
	model "peterchu999/simple-api/model"
	. "peterchu999/simple-api/server"
)

func main() {
	server := SetupServer()
	model.ConnectSqliteDatabase()
	server.Run()
}
