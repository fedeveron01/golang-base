package infrastructure

import (
	"fmt"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	material_usecase "github.com/fedeveron01/golang-base/cmd/core/usecases/material"
	"github.com/fedeveron01/golang-base/cmd/entrypoints"
	material_handler "github.com/fedeveron01/golang-base/cmd/entrypoints/handlers/material"
	handler_subscriptions "github.com/fedeveron01/golang-base/cmd/entrypoints/handlers/subscriptions"
	"github.com/fedeveron01/golang-base/cmd/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//inject dependencies..

type HandlerContainer struct {
	CalculateAge       entrypoints.Handler
	EditSubscription   entrypoints.Handler
	DeleteSubscription entrypoints.Handler
	CreateSubscription entrypoints.Handler
	GetAllMaterial     entrypoints.Handler
	CreateMaterial     entrypoints.Handler
}

func Start() HandlerContainer {

	// inject mysql and gorm
	dsn := "root:root@tcp(127.0.0.1:3306)/fabrica?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(
		entities.Material{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
	fmt.Println("OK")

	// inject repositories

	subscriptionRepository := repositories.NewSubscriptionRepository(db)
	materialRepository := repositories.NewMaterialRepository(db)

	// inject use cases
	materialUseCase := material_usecase.NewMaterialUsecase(materialRepository)

	// inject handlers
	handlerContainer := HandlerContainer{}
	handlerContainer.EditSubscription = handler_subscriptions.NewEditSubscriptionHandler(*subscriptionRepository)
	handlerContainer.DeleteSubscription = handler_subscriptions.NewDeleteSubscriptionHandler(*subscriptionRepository)
	handlerContainer.CreateSubscription = handler_subscriptions.NewCreateSubscriptionHandler(*subscriptionRepository)

	handlerContainer.CreateMaterial = material_handler.NewCreateMaterialHandler(*materialUseCase)
	handlerContainer.GetAllMaterial = material_handler.NewGetAllMaterialHandler(*materialUseCase)

	//handlerContainer.CalculateAge = calculateAgeHandler
	return handlerContainer

}
