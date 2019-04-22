package main

import (
	"context"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	pb "github.com/naichadouban/gomicroservices/user-service/proto/user"
	"log"
	"os"
	"strings"
)

func main() {
	// Create new greeter client
	client := pb.NewUserServiceClient("shippy.service.user", microclient.DefaultClient)
	// Define our flogs
	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "xuxiaofeng",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "13641537547@163.com",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "123456789",
			},
			cli.StringFlag{
				Name:  "company",
				Usage: "china student",
			},
		),
	)
	// start as service
	service.Init(
		micro.Action(func(c *cli.Context) {

			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			company := c.String("company")
			llog.Tracef("parse command args:%v,%v,%v,%v\n",name,email,password,company)
			// Call our user service
			r, err := client.Create(context.TODO(), &pb.User{
				Name:     name,
				Email:    email,
				Password: password,
				Company:  company,
			})
			if err != nil {
				log.Fatalf("Could not create: %v", err)
			}
			llog.Infof("Created: %s", r.User.Id)
			llog.Println(strings.Repeat("=",30))
			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				llog.Fatalf("Could not list users: %v", err)
			}
			for _, v := range getAll.Users {
				llog.Info(v)
			}

			authRes,err := client.Auth(context.TODO(),&pb.User{
				Email:email,
				Password:password,
			})
			if err != nil {
				log.Fatalf("auth failed: %v", err)
			}
			log.Println("token: ", authRes.Token)

			os.Exit(0)
		}),
	)
	// Run the server
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
