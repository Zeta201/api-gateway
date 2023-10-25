package routes

import (
	"github.com/Zeta201/registration-service/controller"
	"github.com/Zeta201/registration-service/infrastructure"
)

type StudentRoute struct {
	Controller controller.StudentController
	Handler    infrastructure.GinRouter
}

func NewStudentRoute(
	controller controller.StudentController,
	handler infrastructure.GinRouter,
) StudentRoute {
	return StudentRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (route StudentRoute) Setup() {
	student := route.Handler.Gin.Group("/api/v1/registration-service")
	{
		student.GET("/", route.Controller.GetStudents)
		student.POST("/register", route.Controller.RegisterStudent)
		student.GET("/:id", route.Controller.GetStudent)
		student.DELETE("/:id", route.Controller.DeleteStudent)
	}
}
