package user_model

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	UserID         uint
	StreakDays     int    `gorm:"column:streak_days" json:"streak_days"`
	TotalDays      int    `gorm:"column:total_days" json:"total_days"`
	GoLastChapter  string `gorm:"column:go_last_chapter" json:"go_last_chapter"`
	CLastChapter   string `gorm:"column:c_last_chapter" json:"c_last_chapter"`
	CppLastChapter string `gorm:"column:cpp_last_chapter" json:"cpp_last_chapter"`
}

func (u *UserInfo) TableName() string {
	return "user_info"
}
