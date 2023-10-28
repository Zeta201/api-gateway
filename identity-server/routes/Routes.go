package routes

import (
	"github.com/Zeta201/identity-server/controller"
	"github.com/Zeta201/identity-server/infrastructure"
)

type UserRoute struct {
	Controller controller.UserController
	Handler    infrastructure.GinRouter
}

func NewUserRoute(
	controller controller.UserController,
	handler infrastructure.GinRouter,
) UserRoute {
	return UserRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (route UserRoute) Setup() {
	user := route.Handler.Gin.Group("/api/v1/identity-service")
	{
		// user.GET("/", route.Controller.GetStudents)
		user.POST("/signup", route.Controller.SignUp)
		user.POST("/login", route.Controller.Login)

		// user.GET("/:id", route.Controller.GetStudent)
		// user.DELETE("/:id", route.Controller.DeleteStudent)
	}
}
