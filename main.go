package main

import (
	. "simple-setup/server"
)

func main() {
	server := SetupServer()
	server.Run()
}
