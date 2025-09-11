package util

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const SecretKey = "secret"

type JWTToken struct {
	Id                   int64  `json:"id"`
	Username             string `json:"username" binding:"omitempty"`
	Email                string `json:"email" binding:"omitempty"`
	jwt.RegisteredClaims        //嵌入jwt的注册声明(包含)
}

type NeedInfo struct {
	Id       int64  `json:"id"`
	Username string `json:"username" binding:"omitempty"`
	Email    string `json:"email" binding:"omitempty"`
}

// 生成token
func GenerateToken(jwtInfo NeedInfo) (string, error) {
	//1.定义Claims（载荷）
	claims := JWTToken{
		Id:       jwtInfo.Id,
		Username: jwtInfo.Username,
		Email:    jwtInfo.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    jwtInfo.Username,                                   //发行者
			Subject:   "token验证",                                          //设置主题
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), //设置过期时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     //设置生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     //设置发行时间
		},
	}
	//2.创建token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //NewWithClaims()是初始化token结构体header并设置加密算法的
	//3.使用秘钥签名并生成完整的token字符串
	tokenString, err := token.SignedString([]byte(SecretKey)) //secret是设置的秘钥名，解析也需要用这个名字
	if err != nil {
		return "", err
	}
	return tokenString, nil //返回token字符串
}
