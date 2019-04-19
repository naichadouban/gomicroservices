package main

import (
	"log"
	"os"

	micro "github.com/micro/go-micro"
	pb "github.com/naichadouban/gomicroservices/consignment-service/proto/consignment"
	vesselProto "github.com/naichadouban/gomicroservices/vessel-service/proto/vessel"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	// 获取容器设置的数据库地址环境变量
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = defaultHost
	}
	session, err := CreateSession(dbHost)
	if err != nil {
		log.Fatalf("create session error: %v\n", err)
	}
	defer session.Close()
	// set-up micro instance
	server := micro.NewService(
		micro.Name("shippy.service.consignment"),
		micro.Version("latest"),
	)
	// Init will parse the command line flags
	server.Init()
	// vessel-service 的客户端
	vesselClient := vesselProto.NewVesselServiceClient("shippy.service.vessel", server.Client())

	// register handler
	pb.RegisterShippingServiceHandler(server.Server(), &handler{session,vesselClient})
	// Run the server
	if err := server.Run(); err != nil {
		log.Panicf("micro server run error:%v\n", err)
	}
}
