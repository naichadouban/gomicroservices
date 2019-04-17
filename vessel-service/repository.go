package main

import (
	"errors"

	"gopkg.in/mgo.v2"

	pb "github.com/naichadouban/gomicroservices/vessel-service/proto/vessel"
)

const (
	DB_NAME        = "shippy"
	CON_COLLECTION = "vessel"
)

type repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
}
type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity < vessel.Capacity && spec.MaxWeight < vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("no vessel find by that spec")
}
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}
func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(DB_NAME).C(CON_COLLECTION)
}
