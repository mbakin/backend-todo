package main

import (
	"backend_todo/server"
	"log"
)

func main() {
	svr := server.NewServer()
	err := svr.StartServer(3000)
	if err != nil {
		log.Fatalln(err)
	}
}
