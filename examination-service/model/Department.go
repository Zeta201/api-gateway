package model

type Department struct {
	ID       int64     `gorm:"primary_key;auto_increment" json:"id"`
	Name     string    `gorm:"size:200" json:"name"`
	Subjects []Subject `gorm:"many2many:dep_subject"`
}

func (department *Department) TableName() string {
	return "Department"
}

// func (subject *Subject) ResponseMap() map[string]interface{} {
// 	resp := make(map[string]interface{})
// 	resp["id"] = subject.ID
// 	resp["name"] = subject.Name
// 	resp["credits"] = subject.Credits
// 	resp["department"] = subject.Departments
// 	return resp
// }
