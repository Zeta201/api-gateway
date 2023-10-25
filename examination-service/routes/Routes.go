package routes

import (
	"github.com/Zeta201/examination-service/controller"
	"github.com/Zeta201/examination-service/infrastructure"
)

type SubjectRoute struct {
	Controller controller.SubjectController
	Handler    infrastructure.GinRouter
}

func NewSubjectRoute(
	controller controller.SubjectController,
	handler infrastructure.GinRouter,
) SubjectRoute {
	return SubjectRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (route SubjectRoute) Setup() {
	subject := route.Handler.Gin.Group("/api/v1/examination-service")
	{
		subject.GET("/", route.Controller.GetAllSubjects)
		subject.POST("/", route.Controller.AddSubject)
		subject.GET("/:id", route.Controller.GetSubjectById)
		// subject.DELETE("/:id", route.Controller.DeleteStudent)
	}
}
