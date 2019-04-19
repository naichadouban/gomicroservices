package main

import (
	"github.com/micro/go-micro"
	"fmt"
	pb "github.com/naichadouban/gomicroservices/user-service/proto/user"
	"log"
)

func main(){
	// 连接到数据库
	db,err := CreateConnection()
	if err!= nil{
		log.Fatalf("connect error: %v\n", err)
	}
	fmt.Printf("%+v\n",db)
	repo := &UserRepository{db:db}
	// 自动检测User结构是否变化了
	db.AutoMigrate(&pb.User{})

	s := micro.NewService(micro.Name("shippy.service.user"),micro.Version("latest"))
	s.Init()

	pb.RegisterUserServiceHandler(s.Server(),&handler{repo:repo})

	if err := s.Run();err!= nil{
		log.Fatalf("user service error: %v\n", err)
	}

}
