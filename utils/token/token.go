package token

import (
	mError "ganxue-server/utils/error"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

// GenerateShortToken 生成短token
func GenerateShortToken(userID uint) (string, *mError.Error) {
	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Minute).Unix(),
	})

	// 生成token字符串
	tokenString, err := token.SignedString([]byte(os.ExpandEnv("$JWT_SECRET")))
	if err != nil {
		return "", mError.New(mError.GenerateTokenError, err, "生成token失败")
	}
	return tokenString, nil
}

// GenerateLongToken 生成长token
func GenerateLongToken(userID uint) (string, *mError.Error) {
	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	// 生成token字符串
	tokenString, err := token.SignedString([]byte(os.ExpandEnv("$JWT_SECRET")))
	if err != nil {
		return "", mError.New(mError.GenerateTokenError, err, "生成token失败")
	}
	return tokenString, nil
}

// ParseToken 解析token
func ParseToken(tokenString string) (uint, *mError.Error) {
	// 解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.ExpandEnv("$JWT_SECRET")), nil
	})
	if err != nil {
		return 0, mError.New(mError.ParseTokenError, err, "解析token失败")
	}

	// 获取token中的claims信息
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, mError.New(mError.ParseTokenError, err, "解析token失败")
	}
	return uint(claims["userID"].(float64)), nil
}
