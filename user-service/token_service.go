package main

import (
	"github.com/dgrijalva/jwt-go"
	pb "github.com/naichadouban/gomicroservices/user-service/proto/user"
	"time"
)

type Authable interface {
	Decode(tokenStr string)(*CustomClaims,error)
	Encode(user *pb.User)(string,error)
}
// 定义加盐hash密码所有的盐
var privateKey = []byte("microsalt")
// 自定义的 metadata，在加密后作为 JWT 的第二部分返回给客户端
type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type TokenService struct {
	repo repository
}
// 将jwt字符串解析成CustomClaims对象
func (t *TokenService) Decode(tokenStr string)(*CustomClaims,error){
	llog.Tracef("Decode tokenStr %s to CustomClaims")
	token,err:=jwt.ParseWithClaims(tokenStr,&CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey,nil
	})
	// 解密转换类型并返回
	if claim,ok := token.Claims.(*CustomClaims);ok && token.Valid{
		return claim,nil
	}else{
		llog.Errorf("jwt parse tokenStr error:%v",err)
		return nil,err
	}
}

// 将 User对象转化为token字符串
func (t *TokenService) Encode(u *pb.User)(string,error){
	// 三天后过期
	expireTime := time.Now().Add(time.Hour*24*3).Unix()
	claim := CustomClaims{
		u,
		jwt.StandardClaims{
			ExpiresAt:expireTime,
			Issuer:"shippy.service.user",
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	return jwtToken.SignedString(privateKey)
}