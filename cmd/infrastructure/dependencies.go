package infrastructure

import (
	"context"
	"fmt"
	"log"

	"github.com/fedeveron01/golang-base/cmd/core/usecases/calculate_age"
	"github.com/fedeveron01/golang-base/cmd/entrypoints"
	handler_person "github.com/fedeveron01/golang-base/cmd/entrypoints/handlers/person"
	handler_subscriptions "github.com/fedeveron01/golang-base/cmd/entrypoints/handlers/subscriptions"
	handler_whatsapp "github.com/fedeveron01/golang-base/cmd/entrypoints/handlers/whatsapp"
	"github.com/fedeveron01/golang-base/cmd/repositories"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//inject dependencies..

type HandlerContainer struct {
	CalculateAge       entrypoints.Handler
	Whatsapp           entrypoints.Handler
	GetAllSubscription entrypoints.Handler
	GetSubscription    entrypoints.Handler
	EditSubscription   entrypoints.Handler
	DeleteSubscription entrypoints.Handler
	CreateSubscription entrypoints.Handler
}

func Start() HandlerContainer {

	// inject mongo db
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://fedeveron2:JtHpIis2IvIVlwLB@sysgestion.jv83i.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Seleccionar la base de datos y crear una instancia del repositorio
	database := client.Database("SysGestion")

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
	whatClient := whatsmeow.NewClient(deviceStore, clientLog)

	// inject repositories

	subscriptionRepository := repositories.NewSubscriptionRepository(database, "subscriptions")
	fmt.Println(subscriptionRepository)

	whatsappRepository := repositories.NewWhatsappRepository(whatClient)

	// inject use cases
	calculateAgeUseCase := calculate_age.Implementation{}

	// inject handlers
	handlerContainer := HandlerContainer{}
	handlerContainer.CalculateAge = handler_person.NewPersonGetAllHandler(calculateAgeUseCase)
	handlerContainer.Whatsapp = handler_whatsapp.NewWhatsappHandlerHandler(*whatsappRepository)
	handlerContainer.GetAllSubscription = handler_subscriptions.NewGetAllSubscriptionHandler(*subscriptionRepository)
	handlerContainer.GetSubscription = handler_subscriptions.NewGetSubscriptionHandler(*subscriptionRepository)
	handlerContainer.EditSubscription = handler_subscriptions.NewEditSubscriptionHandler(*subscriptionRepository)
	handlerContainer.DeleteSubscription = handler_subscriptions.NewDeleteSubscriptionHandler(*subscriptionRepository)
	handlerContainer.CreateSubscription = handler_subscriptions.NewCreateSubscriptionHandler(*subscriptionRepository)

	//handlerContainer.CalculateAge = calculateAgeHandler
	return handlerContainer

}
