package main

import (
	"github.com/Zeta201/identity-server/controller"
	"github.com/Zeta201/identity-server/infrastructure"
	"github.com/Zeta201/identity-server/model"
	"github.com/Zeta201/identity-server/repository"
	"github.com/Zeta201/identity-server/routes"
	"github.com/Zeta201/identity-server/service"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {
	router := infrastructure.NewGinRouter()
	db := infrastructure.NewDatabase()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()
	db.DB.AutoMigrate(&model.User{})
	router.Gin.Run(":7879")
}
