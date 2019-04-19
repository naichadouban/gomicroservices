package main

import (
	"fmt"
	"log"
	"os"

	micro "github.com/micro/go-micro"

	pb "github.com/naichadouban/gomicroservices/vessel-service/proto/vessel"
)

const (
	DEFAULT_HOST = "localhost:27017"
)
func main() {
	host:= os.Getenv("DB_HOST")
	if host==""{
		host = DEFAULT_HOST
	}
	session,err := CreateSession(host)
	defer session.Close()
	if err != nil{
		log.Fatalf("create session error:%v\n",err)
	}
	repo := &VesselRepository{session:session.Copy()}
	CreateDummyData(repo)

	server := micro.NewService(
		micro.Name("shippy.service.vessel"),
		micro.Version("latest"),
	)
	server.Init()
	pb.RegisterVesselServiceHandler(server.Server(), &handler{session})

	if err := server.Run(); err != nil {
		fmt.Println(err)
	}
}
func CreateDummyData(repo repository){
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}
	for _,v := range vessels{
		repo.Create(v)
	}
}
