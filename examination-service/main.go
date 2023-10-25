package main

import (
	"github.com/Zeta201/examination-service/controller"
	"github.com/Zeta201/examination-service/infrastructure"
	"github.com/Zeta201/examination-service/model"
	"github.com/Zeta201/examination-service/repository"
	"github.com/Zeta201/examination-service/routes"
	"github.com/Zeta201/examination-service/service"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {
	router := infrastructure.NewGinRouter()
	db := infrastructure.NewDatabase()
	subjectRepository := repository.NewSubjectRepository(db)
	subjectService := service.NewSubjectService(subjectRepository)
	subjectController := controller.NewSubjectController(subjectService)
	subjectRoute := routes.NewSubjectRoute(subjectController, router)
	subjectRoute.Setup()

	db.DB.AutoMigrate(&model.Department{}, &model.Subject{})
	router.Gin.Run(":9090")

}
