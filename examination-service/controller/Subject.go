package controller

import (
	"net/http"
	"strconv"

	"github.com/Zeta201/examination-service/model"
	"github.com/Zeta201/examination-service/service"
	"github.com/Zeta201/examination-service/util"
	"github.com/gin-gonic/gin"
)

type SubjectController struct {
	service service.SubjectService
}

func NewSubjectController(s service.SubjectService) SubjectController {
	return SubjectController{
		service: s,
	}
}

func (controller SubjectController) GetAllSubjects(ctx *gin.Context) {
	var subjects model.Subject

	keyword := ctx.Query("keyword")

	data, total, err := controller.service.FindAll(subjects, keyword)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find subjects")
		return
	}

	respArr := make([]map[string]interface{}, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Subject result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		},
	})
}

func (controller *SubjectController) AddSubject(ctx *gin.Context) {
	var subjectRequest util.AddSubjectRequest

	ctx.ShouldBindJSON(&subjectRequest)
	// name credits department
	if subjectRequest.Name == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Subject name is required")
	}

	if subjectRequest.Credits <= 0 {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Credits is required")
	}

	if len(subjectRequest.Departments) == 0 {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Subject must be offered at least one department")
	}

	mappedSubject := util.MapToSubject(subjectRequest)

	// var department []model.Department

	// for _,deptID := range subject.Department{

	// 	var dept model.Department
	// 	dept.ID = deptID
	// 	foundDept, err :=
	// 	department = append(department, dept)
	// 	if err != nil {
	// 		util.ErrorJSON(ctx, http.StatusBadRequest, "Error finding student")
	// 		return
	// 	}
	// }

	err := controller.service.SaveSubject(mappedSubject)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to save student")
		return
	}

	util.SuccessJSON(ctx, http.StatusCreated, "Registration Successful")
}

func (controller *SubjectController) GetSubjectById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Invalid id")
		return
	}

	var subject model.Subject
	subject.ID = id
	foundSubject, err := controller.service.GetSubjectById(subject)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error finding subject")
		return
	}

	response := foundSubject.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Subject found",
		Data:    &response,
	})
}
