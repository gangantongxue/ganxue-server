package password

import (
	"ganxue-server/utils/error"
	"golang.org/x/crypto/bcrypt"
)

// EncryptPassword 密码加密
func EncryptPassword(password string) (string, *error.Error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", error.New(error.
			EncryptPasswordError, err, "密码加密失败")
	}
	return string(hashedPwd), nil
}

// ComparePasswords 验证密码
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePlain := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		return false
	}
	return true
}
