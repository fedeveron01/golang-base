package infrastructure

import (
	"github.com/fedeveron01/golang-base/cmd/core/usecases/calculate_age"
	"github.com/fedeveron01/golang-base/cmd/entrypoints"
	handler_person "github.com/fedeveron01/golang-base/cmd/entrypoints/handlers/person"
	"github.com/fedeveron01/golang-base/cmd/repositories"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

//inject dependencies..

type HandlerContainer struct {
	CalculateAge entrypoints.Handler
}

func Start() HandlerContainer {

	// inject whatsapp db
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
	container, err := sqlstore.New("sqlite3", "file:examplestore.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)

	// inject repositories
	whatsappRepository := repositories.NewWhatsappRepository(client)
	whatsappRepository.SendText("5493516247275", "test")

	// inject use cases
	calculateAgeUseCase := calculate_age.Implementation{}

	// inject handlers
	handlerContainer := HandlerContainer{}
	handlerContainer.CalculateAge = handler_person.NewPersonGetAllHandler(calculateAgeUseCase)

	//handlerContainer.CalculateAge = calculateAgeHandler
	return handlerContainer

}
