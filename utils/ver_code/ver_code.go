package ver_code

import (
	"bytes"
	"ganxue-server/global"
	"ganxue-server/utils/error"
	"gopkg.in/mail.v2"
	"math/rand"
	"os"
	"time"
)

// GetVerCode 获取验证码
func GetVerCode(email string) *error.Error {
	verCode := generateVerCode()
	str, err := global.RDB.Set(global.CTX, email, verCode, 15*time.Minute).Result()
	if err != nil {
		return error.New(error.RedisError, err, "redis error"+str)
	}

	if err := sendVerCode(email, verCode); err != nil {
		return err
	}

	return nil
}

// generateVerCode 生成验证码
func generateVerCode() string {
	var verCode string

	full := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for i := 0; i < 6; i++ {
		verCode += string(full[rand.Intn(len(full))])
	}

	return verCode
}

// sendVerCode 发送验证码
func sendVerCode(email string, verCode string) *error.Error {
	m := mail.NewMessage()

	m.SetHeader("From", global.CONFIG.Email.From)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "验证码")

	htmlContent, err := os.ReadFile("utils/ver_code/ver_code_model.html")
	if err != nil {
		return error.New(error.ReadFileError, err, "读取文件失败")
	}

	htmlContent = bytes.ReplaceAll(htmlContent, []byte("{{.VerCode}}"), []byte(verCode))
	m.SetBody("text/html", string(htmlContent))

	// 发送验证码
	d := mail.NewDialer(
		global.CONFIG.Email.Host,
		global.CONFIG.Email.Port,
		global.CONFIG.Email.Username,
		global.CONFIG.Email.Password,
	)
	d.SSL = true

	if err := d.DialAndSend(m); err != nil {
		return error.New(error.SendEmailError, err, "发送邮件失败")
	}

	return nil
}

// Verify 验证验证码
func Verify(email string, verCode string) bool {
	val, err := global.RDB.GetDel(global.CTX, email).Result()
	if err != nil {
		return false
	}

	return val == verCode
}
