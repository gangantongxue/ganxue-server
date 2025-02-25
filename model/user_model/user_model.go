package user_model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//type Model struct {
	//	ID        uint `gorm:"primarykey"`
	//	CreatedAt time.Time
	//	UpdatedAt time.Time
	//	DeletedAt DeletedAt `gorm:"index"`
	//}
	UserName string `gorm:"column: user_name" ,json:"user_name"`
	Password string `gorm:"column: password" ,json:"password"`
	Email    string `gorm:"column: email,uniqueIndex" ,json:"email"`
}

func (u *User) TableName() string {
	return "user"
}
