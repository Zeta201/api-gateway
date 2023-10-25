package model

import "time"

type Student struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Firstname string    `gorm:"size:200" json:"firstname"`
	Lastname  string    `gorm:"size:200" json:"lastname"`
	City      string    `gorm:"size:200" json:"city"`
	Course    string    `gorm:"size:200" json:"course"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (student *Student) TableName() string {
	return "Student"
}

func (student *Student) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = student.ID
	resp["title"] = student.Firstname
	resp["lastname"] = student.Lastname
	resp["city"] = student.City
	resp["course"] = student.Course
	resp["created_at"] = student.CreatedAt
	return resp
}
