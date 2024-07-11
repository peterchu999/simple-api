package main

import (
	. "peterchu999/simple-api/server"
)

func main() {
	server := SetupServer()

	server.Run()
}
