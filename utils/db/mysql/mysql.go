package mysql

import (
	"errors"
	"fmt"
	"ganxue-server/global"
	"ganxue-server/model/user_model"
	"ganxue-server/utils/error"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func Init() {
	// 初始化数据库
	if err := InitDB(); err != nil {
		fmt.Println("mysql init error", err)
		os.Exit(1)
	}
}

func Close() {

}

// InitDB 初始化数据库
func InitDB() *error.Error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.CONFIG.Mysql.User,
		global.CONFIG.Mysql.Password,
		global.CONFIG.Mysql.Host,
		global.CONFIG.Mysql.Port,
		global.CONFIG.Mysql.Database,
	)

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Warn, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return error.New(error.DBConnectError, err, "连接数据库失败")
	}
	global.DB = db

	// 连接池设置
	sqlDB, _err := db.DB()
	if _err != nil {
		return error.New(error.DBConnectError, _err, "连接数据库失败")
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 10)

	// 自动迁移
	if err := global.DB.AutoMigrate(&user_model.User{}, &user_model.UserInfo{}); err != nil {
		return error.New(error.AutoMigrateError, err, "自动迁移失败")
	}

	return nil
}

// Create 创建用户
func Create(user *user_model.User) *error.Error {
	if err := global.DB.Create(user).Error; err != nil {
		return error.New(error.CreateNewUserError, err, "创建用户失败")
	}
	user.UserInfo.UserID = user.ID
	if err := global.DB.Create(&user.UserInfo).Error; err != nil {
		return error.New(error.CreateNewUserError, err, "创建用户失败")
	}
	return nil
}

// FindUserByEmail 根据邮箱查找用户
func FindUserByEmail(email string) (*user_model.User, *error.Error) {
	var user user_model.User
	if err := global.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.New(error.GetUserError, err, "用户未找到")
		}
		return nil, error.New(error.GetUserError, err, "查找用户失败")
	}
	return &user, nil
}

// FindUserByID 根据ID查找用户
func FindUserByID(id uint) (*user_model.User, *error.Error) {
	var user user_model.User
	if err := global.DB.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.New(error.GetUserError, err, "用户未找到")
		}
		return nil, error.New(error.GetUserError, err, "查找用户失败")
	}
	return &user, nil
}

// ChangePassword 修改密码  加密后
func ChangePassword(user interface{}, password string) *error.Error {
	if err := global.DB.Model(user).Update("password", password).Error; err != nil {
		return error.New(error.ChangePasswordError, err, "修改密码失败")
	}
	return nil
}

// Update 更新用户信息
func Update(user interface{}) *error.Error {
	if err := global.DB.Save(user).Error; err != nil {
		return error.New(error.UpdateUserError, err, "更新用户失败")
	}
	return nil
}

// Delete 删除用户
func Delete(user interface{}) *error.Error {
	if err := global.DB.Delete(user).Error; err != nil {
		return error.New(error.DeleteUserError, err, "删除用户失败")
	}
	return nil
}

// FindUserInfoByID 根据ID查找用户信息
func FindUserInfoByID(id uint) (*user_model.User, *error.Error) {
	var user user_model.User
	if err := global.DB.Preload("UserInfo").Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.New(error.GetUserError, err, "用户未找到")
		}
		return nil, error.New(error.GetUserError, err, "查找用户失败")
	}
	return &user, nil
}
