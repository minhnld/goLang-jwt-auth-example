package main

import (
	"./db"
	"./server"
	"./server/middleware/myJwt"
	"log"
)

var host = "localhost"
var port = "8080"

func main() {
	// init the DB
	db.InitDB()

	// init the JWTs
	jwtErr := myJwt.InitJWT()
	if jwtErr != nil {
		log.Println("Error initializing the JWT's!")
		log.Fatal(jwtErr)
	}

	// start the server
	serverErr := server.StartServer(host, port)
	if serverErr != nil {
		log.Println("Error starting server!")
		log.Fatal(serverErr)
	}
}
