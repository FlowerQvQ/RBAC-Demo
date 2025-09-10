package middleware

import (
	"NewProject/pkg/wapper"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const SecretKey = "secret"

type JWTToken struct {
	Id                   int    `json:"id"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        //嵌入jwt的注册声明
}

// 生成token
func GenerateToken(jwtInfo JWTToken) (string, error) {
	//1.定义Claims（载荷）
	claims := JWTToken{
		Id:       jwtInfo.Id,
		Username: jwtInfo.Username,
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

// 验证token
func ParseToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从请求头重获取token
		tokenString := c.GetHeader("Authorization") //Authorization用来接token，前端传递的时候也要用这个名字传递
		if tokenString == "" {
			wapper.ResError(c, wapper.TokenIsNull)
			c.Abort() //c.Abort()会阻止后续的中间件和处理函数执行，但它不会立即终止当前函数的执行
			return    //return会立即终止当前函数的执行
		}
		//解析token：解析JWT令牌并验证其签名
		//ParseWithClaims()函数会自动判断token是否过期，不需要再写判断语句
		token, err := jwt.ParseWithClaims(tokenString, &JWTToken{}, func(token *jwt.Token) (interface{}, error) {
			//确保方法是HMAC
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				//当解析JWT令牌时，发现其签名方法不是 HMAC ,则服务端会替换当前的方法并返回：
				//eg:例如使用了 RS256 签名方法,会返回："unexpected signing method: RS256"
				//token.Header["alg"]是JWT（JSON Web Token）头部中的一个字段，表示签名算法
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(SecretKey), nil
		})
		//如果令牌解析出错，返回错误响应
		if err != nil {
			wapper.ResError(c, wapper.AuthenticationFailed)
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(*JWTToken); ok {
			c.Set("username", claims.Subject)
			c.Set("userId", claims.ID)
			c.Next()
		} else {
			wapper.ResError(c, wapper.PayLoadParsingFailed)
			c.Abort()
			return
		}
	}
}
