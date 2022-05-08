package main

import (
	"log"

	"pantry/server"
)

func main() {
	app := server.New()

	log.Fatal(app.Listen(":3000"))
}
