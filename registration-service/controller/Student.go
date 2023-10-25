package controller

import (
	"net/http"
	"strconv"

	"github.com/Zeta201/registration-service/model"
	"github.com/Zeta201/registration-service/service"
	"github.com/Zeta201/registration-service/util"
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	service service.StudentService
}

func NewStudentController(s service.StudentService) StudentController {
	return StudentController{
		service: s,
	}
}

func (controller StudentController) GetStudents(ctx *gin.Context) {
	var students model.Student

	keyword := ctx.Query("keyword")

	data, total, err := controller.service.FindAll(students, keyword)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find students")
		return
	}

	respArr := make([]map[string]interface{}, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Student result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		},
	})
}

func (controller *StudentController) RegisterStudent(ctx *gin.Context) {
	var student model.Student

	ctx.ShouldBindJSON(&student)

	if student.Firstname == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Firstname is required")
	}

	if student.Lastname == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Lastname is required")
	}

	if student.City == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Email is required")
	}

	if student.Course == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Course is required")
	}

	err := controller.service.Save(student)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to save student")
		return
	}

	util.SuccessJSON(ctx, http.StatusCreated, "Registration Successful")
}

func (controller *StudentController) GetStudent(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Invalid id")
		return
	}

	var student model.Student
	student.ID = id
	foundStudent, err := controller.service.Find(student)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error finding student")
		return
	}

	response := foundStudent.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Student found",
		Data:    &response,
	})
}

func (controller *StudentController) DeleteStudent(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}

	err = controller.service.Delete(id)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error deleting student")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Student deleted",
	}
	ctx.JSON(http.StatusOK, response)
}
