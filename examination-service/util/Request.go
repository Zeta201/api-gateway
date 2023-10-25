package util

import "github.com/Zeta201/examination-service/model"

type AddSubjectRequest struct {
	ID          int64
	Name        string
	Credits     int32
	Departments []int64
}

func MapToDepartments(departments []int64) []model.Department {

	var deps []model.Department
	for _, dept := range departments {
		var d model.Department
		d.ID = dept
		deps = append(deps, d)
	}
	return deps
}

func MapToSubject(subjectRequest AddSubjectRequest) model.Subject {
	return model.Subject{
		ID:          subjectRequest.ID,
		Name:        subjectRequest.Name,
		Credits:     subjectRequest.Credits,
		Departments: MapToDepartments(subjectRequest.Departments),
	}
}
