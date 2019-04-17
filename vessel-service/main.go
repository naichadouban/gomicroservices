package main

import (
	"context"
	"errors"
	"fmt"

	micro "github.com/micro/go-micro"

	pb "github.com/naichadouban/gomicroservices/vessel-service/proto/vessel"
)

type repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}
type VesselRepository struct {
	vessels []*pb.Vessel
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity < vessel.Capacity && spec.MaxWeight < vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("no vessel find by that spec")
}

type service struct {
	repo repository
}

func (s *service) FindAvailableVessel(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}
	res.Vessel = vessel
	return nil
}

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
