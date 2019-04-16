package main

import (
	"context"
	"log"
	"sync"

	micro "github.com/micro/go-micro"
	pb "github.com/naichadouban/gomicroservices/consignment-service/proto/consignment"
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
	repo *Repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
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
	// register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})
	// Run the server
	if err := srv.Run(); err != nil {
		log.Panicf("micro server run error:%v\n", err)
	}
}
