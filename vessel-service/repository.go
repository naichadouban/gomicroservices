package main

import (
	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"

	pb "github.com/naichadouban/gomicroservices/vessel-service/proto/vessel"
)

const (
	DB_NAME        = "vessels"
	CON_COLLECTION = "vessels"
)

type repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
	Close()
}
type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel
	err := repo.collection().Find(bson.M{
		"capacity":bson.M{"$gte":spec.Capacity},
		"maxweight":bson.M{"$gte":spec.MaxWeight},
	}).One(&vessel)
	if err != nil{
		return nil,err
	}
	return vessel,nil
}
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}
func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(DB_NAME).C(CON_COLLECTION)
}

func (repo *VesselRepository) Close(){
	repo.session.Close()
}