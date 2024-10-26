package infrastructure

import (
	"log"
	"net/http"
	"os"

	"github.com/VsenseTechnologies/skf_plc_http_server/infrastructure/db"
	"github.com/VsenseTechnologies/skf_plc_http_server/infrastructure/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/infrastructure/smtpclient"
	"github.com/VsenseTechnologies/skf_plc_http_server/presentation/route"
)

func Run() {
	serverAddress := os.Getenv("SERVER_ADDRESS")

	if serverAddress == "" {
		log.Fatalln("missing environment variable SERVER_ADDRESS")
	}

	database, error := db.Connect()

	if error != nil {
		log.Fatalln("failed to connect to database: ", error.Error())
	}

	log.Println("connected to database")

	smtpClient := smtpclient.SetupClient()

	postgresRepository := repository.NewPostgresRepository(database)
	smtpClientRepository := repository.NewSmtpClientRepository(&smtpClient)

	router := route.Router(&postgresRepository, &smtpClientRepository)

	log.Printf("server is running on %s", serverAddress)
	http.ListenAndServe(serverAddress, router)

}
