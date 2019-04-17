package main

import (
	"fmt"

	micro "github.com/micro/go-micro"

	pb "github.com/naichadouban/gomicroservices/vessel-service/proto/vessel"
)

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}
	repo := &VesselRepository{vessels}
	srv := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)
	srv.Init()
	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
