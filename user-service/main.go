package main

import (
	"fmt"
	"github.com/micro/go-micro"
	pb "github.com/naichadouban/gomicroservices/user-service/proto/user"
)

func main(){
	// 连接到数据库
	db,err := CreateConnection()
	if err!= nil{
		llog.Fatalf("connect error: %v\n", err)
	}
	fmt.Printf("%+v\n",db)
	repo := &UserRepository{db:db}
	// 自动检测User结构是否变化了
	db.AutoMigrate(&pb.User{})

	s := micro.NewService(micro.Name("shippy.service.user"),micro.Version("latest"))
	s.Init()

	pb.RegisterUserServiceHandler(s.Server(),&handler{
		repo:repo,
		tokenService:&TokenService{},
	})

	if err := s.Run();err!= nil{
		llog.Fatalf("user service error: %v\n", err)
	}
}

