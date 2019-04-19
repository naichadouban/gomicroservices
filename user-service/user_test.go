package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	pb "github.com/naichadouban/gomicroservices/user-service/proto/user"
	"testing"
	"time"
)

func Test(t *testing.T){
	// 三天后过期
	expireTime := time.Now().Add(time.Hour*24*3).Unix()
	claim := CustomClaims{
		&pb.User{
			Name:"xuxiaofeng",
			Password:"123456",
		},
		jwt.StandardClaims{
			ExpiresAt:expireTime,
			Issuer:"shippy.service.user",
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	ts,_ := jwtToken.SignedString(privateKey)
	fmt.Println(ts)
}
