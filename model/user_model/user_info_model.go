package user_model

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	UserID     uint
	StreakDays int `gorm:"column:streak_days" json:"streak_days"`
	TotalDays  int `gorm:"column:total_days" json:"total_days"`
}

func (u *UserInfo) TableName() string {
	return "user_info"
}
