package model

type User struct {
	ID       int64  `gorm:"primary_key;auto_increment" json:"id"`
	Username string `gorm:"size:200" json:"username"`
	Password string `gorm:"size:200" json:"password"`
}

func (user *User) TableName() string {
	return "User"
}

func (user *User) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = user.ID
	resp["username"] = user.Username
	resp["password"] = user.Password
	return resp
}
