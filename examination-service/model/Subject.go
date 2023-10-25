package model

type Subject struct {
	ID          int64        `gorm:"primary_key;auto_increment" json:"id"`
	Name        string       `gorm:"size:200" json:"name"`
	Credits     int32        `gorm:"not null" json:"credits"`
	Departments []Department `gorm:"many2many:dep_subject" json:"departments"`
}

func (subject *Subject) TableName() string {
	return "Subject"
}

func (subject *Subject) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = subject.ID
	resp["name"] = subject.Name
	resp["credits"] = subject.Credits
	resp["department"] = subject.Departments
	return resp
}
