package main

import (
	"backend_todo/server"
	"log"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	
	if port == "" {
		port = "3000"
	}

	svr := server.NewServer()
	err := svr.StartServer(port)
	if err != nil {
		log.Fatalln(err)
	}
}
