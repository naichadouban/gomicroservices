package main

import (
	"context"

	pb "github.com/naichadouban/gomicroservices/vessel-service/proto/vessel"
)

type handler struct {
	session *mongo.Co
}

func(h *handler) GetRepo() repository {
	return &VesselRepository{h.session.Clone()}
}
func (h *handler) Create(ctx context.Context,req *pb.Vessel,res *pb.Response)error{
	defer h.GetRepo().Close()
	if err := h.GetRepo().Create(req);err!= nil{
		return err
	}
	res.Vessel = req
	res.Created =true
	return nil

}

func (h *service) FindAvailableVessel(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}
	res.Vessel = vessel
	return nil
}