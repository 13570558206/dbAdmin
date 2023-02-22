package model

type User struct {
	ID       uint `gorm:"primaryKey"`
	Account  uint
	Mobile   string
	Username string
}

// TableName 定义表名
func (User) TableName() string {
	return "u_user"
}
