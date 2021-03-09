package utils

import (
	"errors"
	"fmt"
	"os"
	"time"
	"github.com/dgrijalva/jwt-go"
)
type ReturnComJson struct {
	Code int "code"
	Msg string "msg"
	Data interface{} "data"
}
type JwtUserInfo struct {
	Other string `json:"other"`
}

type JwtData struct {
	Username string `json:"username"`
	Userid int `json:"userid"`
	Userinfo JwtUserInfo `json:"userinfo"`
	jwt.StandardClaims
}
func GenToken(userid int,username string,other JwtUserInfo) (string, error) {
	// 创建一个我们自己的声明
	d, _ := time.ParseDuration(GetIniVal("tokenexpire","")+"m")
	c := JwtData{
		username, // 自定义字段
		userid,
		other,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(8* d).Unix(), // 过期时间
			Issuer:   GetIniVal("jwtuser",""),                               // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString([]byte(GetIniVal("jwtkey","")))
}
func ParseToken(tokenString string) (*JwtData, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &JwtData{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(GetIniVal("jwtkey","")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtData); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func Mkdir(path string) bool  {
	err := os.MkdirAll(path, 0777)
	if err != nil {
		fmt.Printf("crete fail %s", err)
		return false
	}
	return true
}