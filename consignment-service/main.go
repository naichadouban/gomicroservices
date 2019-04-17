package main

import (
	"context"
	"log"
	"sync"

	micro "github.com/micro/go-micro"
	pb "github.com/naichadouban/gomicroservices/consignment-service/proto/consignment"
	vesselProto "github.com/naichadouban/gomicroservices/vessel-service/proto/vessel"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}
type Repository struct {
	mu           sync.Mutex
	consignments []*pb.Consignment
}

// create a new consignment
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// servie should implement all method to satify the service we defined in our protobuf definition.
// you can check the interface in the generated code itself for exact method signatures etc
// to give you a better idea.
type service struct {
	repo         *Repository
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	// here we call a client instance of our vessel service with our consignment weight ,
	// and the amount of containers as the capacity value
	vesselResponse, err := s.vesselClient.FindAvailableVessel(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if err != nil {
		return err
	}
	log.Printf("Found vessel:%s \n", vesselResponse.Vessel.Name)
	// We set the VesselID as the vessel we got back from out vessel vervice
	req.VesselId = vesselResponse.Vessel.Id

	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	res.Consignment = consignment
	res.Created = true
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &Repository{}
	// create a new service
	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)
	// Init will parse the command line flags
	srv.Init()

	vesselClient := vesselProto.NewVesselServiceClient("shippy.service.vessel", srv.Client())

	// register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})
	// Run the server
	if err := srv.Run(); err != nil {
		log.Panicf("micro server run error:%v\n", err)
	}
}
