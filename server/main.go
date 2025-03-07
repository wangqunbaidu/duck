package main

import (
	"duck/internal/server"
	_ "duck/internal/services/eventhandler"
)

func main() {
	server.Init()
	server.NewServer()
}
