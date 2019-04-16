package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	pb "github.com/naichadouban/gomicroservices/consignment-service/proto/consignment"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:8010"
	defaultFileName = "./consignment.json"
)

func main() {
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Panicf("establish connection error:%v\n", err)
	}
	defer cc.Close()
	client := pb.NewShippingServiceClient(cc)
	fileName := defaultFileName
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}
	consignment, err := parseFile(fileName)
	if err != nil {
		log.Panicf("read json file error:%v\n", err)
	}
	res, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Panicf("grpc client call error:%v\n", err)
	}
	// log.Printf("get response:%#v", res)
	log.Printf("create consignment:%v", res.Created)

	log.Println(strings.Repeat("=", 40))
	getAllRes, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Panicf("call GetConsignments error:%v\n", err)
	}
	for _, v := range getAllRes.Consignments {
		log.Println(v)
	}
}

func parseFile(fileName string) (*pb.Consignment, error) {
	var consignment pb.Consignment
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, err
	}
	return &consignment, nil
}
