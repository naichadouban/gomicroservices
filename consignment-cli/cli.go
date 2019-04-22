package main

import (
	"context"
	"encoding/json"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"io/ioutil"
	"log"
	"os"
	"strings"

	pb "github.com/naichadouban/gomicroservices/consignment-service/proto/consignment"
)

const (
	address         = "localhost:8010"
	defaultFileName = "./consignment.json"
)

func main() {
	// 这个名字 shippy.service.consignment 一定要写对啊
	client := pb.NewShippingServiceClient("shippy.service.consignment", microclient.DefaultClient)


	// Define our flogs
	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "file",
				Usage: "consignment file",
			},
			cli.StringFlag{
				Name:  "token",
				Usage: "user token",
			},
		),
	)
	// start as service
	service.Init(
		micro.Action(func(c *cli.Context) {
			fileName := defaultFileName

			file := c.String("file")
			token := c.String("token")

			llog.Infof("parse command args:%v,%v",file,token)
			if file!= "" {
				fileName = file
			}
			if fileName==""||token==""{
				llog.Warnf("fileName or token is empty")
				os.Exit(0)
			}
			// 解析货物信息
			consignment, err := parseFile(fileName)
			if err != nil {
				log.Panicf("read json file error:%v\n", err)
			}
			// 创建带有用户token的context
			ctxWithToken := metadata.NewContext(context.Background(), map[string]string{
				"token":token,
			})
			// 调用rpc
			res, err := client.CreateConsignment(ctxWithToken, consignment)
			if err != nil {
				llog.Panicf("rpc client call error:%v\n", err)
				os.Exit(0)
			}
			// log.Printf("get response:%#v", res)
			llog.Tracef("create consignment:%v", res.Created)

			llog.Infof(strings.Repeat("=", 40))
			// 列出目前所有托用的货物
			getAllRes, err := client.GetConsignments(ctxWithToken, &pb.GetRequest{})
			if err != nil {
				llog.Panicf("call GetConsignments error:%v\n", err)

			}
			for i,c:= range getAllRes.Consignments {
				llog.Printf("consignment_%d: %v\n", i, c)
			}

			os.Exit(0)
		}),
	)
	// Run the server
	if err := service.Run(); err != nil {
		log.Println(err)
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
