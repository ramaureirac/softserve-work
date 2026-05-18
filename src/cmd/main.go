package main

import (
	"log"

	server "github.com/ramaureirac/softserve-work/src/server"
)

func main() {
	log.Println("Server running in http://localhost:8080")
	srv := server.NewRouterApp()
	srv.Run()
}
