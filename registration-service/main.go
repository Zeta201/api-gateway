package main

import (

	// "github.com/Zeta201/registration-service/infrastructure"

	"github.com/Zeta201/registration-service/controller"
	"github.com/Zeta201/registration-service/infrastructure"
	"github.com/Zeta201/registration-service/model"
	"github.com/Zeta201/registration-service/repository"
	"github.com/Zeta201/registration-service/routes"
	"github.com/Zeta201/registration-service/service"
	_ "github.com/go-sql-driver/mysql"
)

// func checkoutBook(c *gin.Context) {
// 	id, ok := c.GetQuery("id")

// 	if !ok {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{
// 			"message": "Missing id query param",
// 		})
// 		return
// 	}

//		book, err := getBookById(id)
//		if err != nil {
//			c.IndentedJSON(http.StatusNotFound, gin.H{
//				"message": "Book not found",
//			})
//			return
//		}
//		if book.Quantity <= 0 {
//			c.IndentedJSON(http.StatusBadRequest, gin.H{
//				"message": "Book not available",
//			})
//			return
//		}
//		book.Quantity -= 1
//		c.IndentedJSON(http.StatusOK, book)
//	}

func init() {
	infrastructure.LoadEnv()
}

// func checkHealth(c *gin.Context) {

// 	infrastructure.LoadEnv()
// 	infrastructure.NewDatabase()
// 	c.IndentedJSON(http.StatusOK, struct {
// 		Status string
// 	}{
// 		Status: "alive",
// 	})
// }

func main() {

	router := infrastructure.NewGinRouter()
	db := infrastructure.NewDatabase()
	studentRepository := repository.NewStudentRepository(db)
	studentService := service.NewStudentService(studentRepository)
	studentController := controller.NewStudentController(studentService)
	studentRoute := routes.NewStudentRoute(studentController, router)
	studentRoute.Setup()

	db.DB.AutoMigrate(&model.Student{})
	router.Gin.Run(":8080")
}
