package main

import (
	pb "github.com/naichadouban/gomicroservices/consignment-service/proto/consignment"
	"gopkg.in/mgo.v2"
)

const (
	DB_NAME        = "shippy"
	CON_COLLECTION = "consignment"
)

type repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close()
}
type ConsignmentRepository struct {
	session *mgo.Session
}

// create a new consignment
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) error {
	return repo.collection().Insert(consignment)
}

// GetAll get all consignment
func (repo *ConsignmentRepository) GetAll() ([]*pb.Consignment, error) {
	var cons []*pb.Consignment
	err := repo.collection().Find(nil).All(&cons)
	return cons, err
}
func (repo *ConsignmentRepository) Close() {
	repo.session.Close()
}

// DB():切换数据库，C():切换集合
func (repo *ConsignmentRepository) collection() *mgo.Collection {
	return repo.session.DB(DB_NAME).C(CON_COLLECTION)
}
