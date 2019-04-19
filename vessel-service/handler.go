package main

import (
	"context"
	"gopkg.in/mgo.v2"

	pb "github.com/naichadouban/gomicroservices/vessel-service/proto/vessel"
)

type handler struct {
	session *mgo.Session
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

func (h *handler) FindAvailableVessel(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	defer h.GetRepo().Close()
	vessel, err := h.GetRepo().FindAvailable(req)
	if err != nil {
		return err
	}
	res.Vessel = vessel
	return nil
}