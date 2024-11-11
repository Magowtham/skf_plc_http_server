package main

import (
	"log"

	"github.com/VsenseTechnologies/skf_plc_http_server/infrastructure"
	"github.com/joho/godotenv"
)

func main() {
	// enable this in development mode
	error := godotenv.Load()

	if error != nil {
		log.Fatalln("Failed to load environment variables: ", error.Error())
	}

	initLogger()

	infrastructure.Run()
}
