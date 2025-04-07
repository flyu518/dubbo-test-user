package model

// User 定义用户模型
type User struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;type:varchar(255);not null;index"`
	Password string `json:"password" gorm:"column:password;type:varchar(255);not null"`
	Phone    string `json:"phone" gorm:"column:phone;type:varchar(255);not null;index"`
	Email    string `json:"email" gorm:"column:email;type:varchar(255);index"`
}

// TableComment 设置表 comment
func (m *User) TableComment() string {
	return "用户表"
}
