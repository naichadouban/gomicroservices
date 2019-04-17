package main

import (
	"context"
	"log"

	"gopkg.in/mgo.v2"

	pb "github.com/naichadouban/gomicroservices/consignment-service/proto/consignment"
	vesselPb "github.com/naichadouban/gomicroservices/vessel-service/proto/vessel"
)

type handler struct {
	session      *mgo.Session
	vesselClient vesselPb.VesselServiceClient
}

func (h *handler) GetRepo() repository {
	return &ConsignmentRepository{h.session.Clone()}
}

func (h *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	defer h.GetRepo().Close()

	// 寻找合适的 轮船
	vesselReq := &vesselPb.Specification{
		Capacity:  int32(len(req.Containers)),
		MaxWeight: req.Weight,
	}

	// here we call a client instance of our vessel service with our consignment weight ,
	// and the amount of containers as the capacity value
	vesselRes, err := h.vesselClient.FindAvailableVessel(context.Background(), vesselReq)
	if err != nil {
		return err
	}
	log.Printf("Found vessel:%s \n", vesselRes.Vessel.Name)
	// We set the VesselID as the vessel we got back from out vessel vervice
	req.VesselId = vesselRes.Vessel.Id

	// Save our consignment
	if err := h.GetRepo().Create(req); err != nil {
		return err
	}
	res.Consignment = req
	res.Created = true
	return nil
}

func (h *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	defer h.GetRepo().Close()
	consignments, err := h.GetRepo().GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
