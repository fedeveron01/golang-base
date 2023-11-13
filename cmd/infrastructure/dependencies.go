package infrastructure

import (
	"fmt"

	material_type_handler "github.com/fedeveron01/golang-base/cmd/adapters/entrypoints/handlers/material_type"
	movement_handler "github.com/fedeveron01/golang-base/cmd/adapters/entrypoints/handlers/movement"
	product_handler "github.com/fedeveron01/golang-base/cmd/adapters/entrypoints/handlers/product"
	material_type_usecase "github.com/fedeveron01/golang-base/cmd/usecases/material_type"
	movement_usecase "github.com/fedeveron01/golang-base/cmd/usecases/movement"
	product_usecase "github.com/fedeveron01/golang-base/cmd/usecases/product"

	"github.com/fedeveron01/golang-base/cmd/adapters/entrypoints"
	charge_handler "github.com/fedeveron01/golang-base/cmd/adapters/entrypoints/handlers/charge"
	employee_handler "github.com/fedeveron01/golang-base/cmd/adapters/entrypoints/handlers/employee"
	material_handler "github.com/fedeveron01/golang-base/cmd/adapters/entrypoints/handlers/material"
	ping_handler "github.com/fedeveron01/golang-base/cmd/adapters/entrypoints/handlers/ping"
	user_handler "github.com/fedeveron01/golang-base/cmd/adapters/entrypoints/handlers/user"
	"github.com/fedeveron01/golang-base/cmd/adapters/gateways"
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/repositories"
	charge_usecase "github.com/fedeveron01/golang-base/cmd/usecases/charge"
	employee_usecase "github.com/fedeveron01/golang-base/cmd/usecases/employee"
	material_usecase "github.com/fedeveron01/golang-base/cmd/usecases/material"
	user_usecase "github.com/fedeveron01/golang-base/cmd/usecases/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//inject dependencies..

type HandlerContainer struct {
	Ping entrypoints.Handler

	MaterialHandler     material_handler.MaterialHandlerInterface
	MaterialTypeHandler material_type_handler.MaterialTypeHandlerInterface
	UserHandler         user_handler.UserHandlerInterface
	ChargeHandler       charge_handler.ChargeHandlerInterface
	EmployeeHandler     employee_handler.EmployeeHandlerInterface
	ProductHandler      product_handler.ProductHandlerInterface
	MovementHandler     movement_handler.MovementHandlerInterface
}

func Start() HandlerContainer {
	// inject mysql and gorm
	dsn := "admin:software-factory-db12@tcp(20.55.53.24:3306)/factory?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(
		gateway_entities.User{}, gateway_entities.Charge{}, gateway_entities.Employee{}, gateway_entities.Material{},
		gateway_entities.MaterialProduct{}, gateway_entities.MaterialType{},
		gateway_entities.Product{}, gateway_entities.Movement{}, gateway_entities.MovementDetail{}, gateway_entities.ProductVariation{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
	fmt.Println("OK")

	// inject repositories
	materialRepository := repositories.NewMaterialRepository(db)
	sessionRepository := repositories.NewSessionRepository(db)
	employeeRepository := repositories.NewEmployeeRepository(db)
	chargeRepository := repositories.NewChargeRepository(db)
	userRepository := repositories.NewUserRepository(db)
	materialTypeRepository := repositories.NewMaterialTypeRepository(db)
	productRepository := repositories.NewProductRepository(db)
	movementRepository := repositories.NewMovementRepository(db)
	movementDetailRepository := repositories.NewMovementDetailRepository(db)

	// inject gateways
	materialGateway := gateways.NewMaterialGateway(materialRepository)
	userGateway := gateways.NewUserGateway(userRepository)
	sessionGateway := gateways.NewSessionGateway(sessionRepository)
	employeeGateway := gateways.NewEmployeeGateway(employeeRepository)
	chargeGateway := gateways.NewChargeGateway(chargeRepository)
	materialTypeGateway := gateways.NewMaterialTypeGateway(materialTypeRepository)
	productGateway := gateways.NewProductGateway(productRepository)
	movementGateway := gateways.NewMovementGateway(movementRepository)
	movementDetailGateway := gateways.NewMovementDetailGateway(movementDetailRepository)

	// inject use cases
	materialUseCase := material_usecase.NewMaterialUsecase(materialGateway, materialTypeGateway)
	userUseCase := user_usecase.NewUserUseCase(userGateway, sessionGateway, employeeGateway, chargeGateway)
	chargeUseCase := charge_usecase.NewChargeUsecase(chargeGateway)
	employeeUseCase := employee_usecase.NewEmployeeUseCase(employeeGateway, chargeGateway)
	materialTypeUseCase := material_type_usecase.NewMaterialTypeUsecase(materialTypeGateway)
	productUseCase := product_usecase.NewProductUsecase(productGateway, materialGateway)
	movementUseCase := movement_usecase.NewMovementUseCase(movementGateway, movementDetailGateway, materialGateway)

	// inject handlers
	handlerContainer := HandlerContainer{}

	handlerContainer.Ping = ping_handler.NewPingHandler()
	handlerContainer.MaterialHandler = material_handler.NewMaterialHandler(sessionGateway, materialUseCase)
	handlerContainer.UserHandler = user_handler.NewUserHandler(sessionGateway, userUseCase)
	handlerContainer.ChargeHandler = charge_handler.NewChargeHandler(sessionGateway, chargeUseCase)
	handlerContainer.EmployeeHandler = employee_handler.NewEmployeeHandler(sessionGateway, employeeUseCase)
	handlerContainer.MaterialTypeHandler = material_type_handler.NewMaterialTypeHandler(sessionGateway, materialTypeUseCase)
	handlerContainer.ProductHandler = product_handler.NewProductHandler(sessionGateway, productUseCase)
	handlerContainer.MovementHandler = movement_handler.NewMovementHandler(sessionGateway, movementUseCase)

	return handlerContainer

}
