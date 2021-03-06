package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"gowebbase/modules/config"
	"math/rand"
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
	Role []int `json:"role"`
}

type JwtData struct {
	Username string `json:"username"`
	Userid int `json:"userid"`
	Userinfo JwtUserInfo `json:"userinfo"`
	jwt.StandardClaims
}
func GenToken(userid int,username string,other JwtUserInfo) (string, error) {
	// 创建一个我们自己的声明
	d, _ := time.ParseDuration(config.GetIniVal("tokenexpire","")+"m")
	c := JwtData{
		username, // 自定义字段
		userid,
		other,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(8* d).Unix(), // 过期时间
			Issuer:   config.GetIniVal("jwtuser",""),                               // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString([]byte(config.GetIniVal("jwtkey","")))
}
func ParseToken(tokenString string) (*JwtData, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &JwtData{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(config.GetIniVal("jwtkey","")), nil
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

// MD5 生成MD5摘要
func MD5(s string) string {
	data := []byte(s)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return  md5str1
}

// MD5 byte生成MD5摘要
func MD5bt(s []byte) string {
	m := md5.New()
	m.Write(s)
	return hex.EncodeToString(m.Sum(nil))
}

func CreateCaptcha6() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
