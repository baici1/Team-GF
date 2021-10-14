package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//定义错误信息
var (
	tokenInvalid = errors.New("token已失效")
)

// MyCustomClaims token得负载
type MyCustomClaims struct {
	ID int64
	jwt.StandardClaims
}

// TokenExpireDuration 定义过期使时间(为一天)
const TokenExpireDuration = time.Hour * 24 * 365

//定义指定的secret签名
var mysecret = []byte("杨澳宇最帅！")

// GenerateToken 生成token
func GenerateToken(id int64) (string, error) {
	//创建一个自定义的负载
	claims := MyCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //过期时间
			Issuer:    "YAY",
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(mysecret)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mysecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, tokenInvalid
}
