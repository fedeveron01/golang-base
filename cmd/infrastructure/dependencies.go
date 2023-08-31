package infrastructure

import (
	"fmt"
	entities2 "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	material_usecase "github.com/fedeveron01/golang-base/cmd/core/usecases/material"
	user_usecase "github.com/fedeveron01/golang-base/cmd/core/usecases/user"
	"github.com/fedeveron01/golang-base/cmd/entrypoints"
	material_handler "github.com/fedeveron01/golang-base/cmd/entrypoints/handlers/material"
	user_handler "github.com/fedeveron01/golang-base/cmd/entrypoints/handlers/user"
	"github.com/fedeveron01/golang-base/cmd/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//inject dependencies..

type HandlerContainer struct {
	//material
	GetAllMaterial entrypoints.Handler
	CreateMaterial entrypoints.Handler
	//user
	CreateUser entrypoints.Handler
	LoginUser  entrypoints.Handler
}

func Start() HandlerContainer {

	// inject mysql and gorm
	dsn := "root:root@tcp(127.0.0.1:3306)/fabrica?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(
		entities2.User{}, entities2.Charge{}, entities2.Employee{}, entities2.Material{},
		entities2.MaterialProduct{}, entities2.MaterialType{}, entities2.MeasurementUnit{},
		entities2.Product{}, entities2.ProductionOrder{}, entities2.ProductionOrderDetail{},
		entities2.PurchaseOrder{}, entities2.PurchaseOrderDetail{}, entities2.Session{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
	fmt.Println("OK")

	// inject repositories
	materialRepository := repositories.NewMaterialRepository(db)
	userRepository := repositories.NewUserRepository(db)
	sessionRepository := repositories.NewSessionRepository(db)
	employeeRepository := repositories.NewEmployeeRepository(db)

	// inject use cases
	materialUseCase := material_usecase.NewMaterialUsecase(materialRepository)
	userUseCase := user_usecase.NewUserUsecase(userRepository, sessionRepository, employeeRepository)

	// inject handlers
	handlerContainer := HandlerContainer{}

	handlerContainer.CreateMaterial = material_handler.NewCreateMaterialHandler(*materialUseCase)
	handlerContainer.GetAllMaterial = material_handler.NewGetAllMaterialHandler(*materialUseCase)

	handlerContainer.CreateUser = user_handler.NewCreateUserHandler(*userUseCase)
	handlerContainer.LoginUser = user_handler.NewLoginUserHandler(*userUseCase)

	//handlerContainer.CalculateAge = calculateAgeHandler
	return handlerContainer

}
