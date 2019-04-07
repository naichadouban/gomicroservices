package main

import (
	"fmt"
	pb "github.com/naichadouban/gomicroservices/consignment-service/proto/consignment"
)

const (
	port = "50051"
)

type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

func main() {

}
