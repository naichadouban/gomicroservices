package main

import (
	"context"
	"errors"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	userProto "github.com/naichadouban/gomicroservices/user-service/proto/user"
	"log"
	"os"

	"github.com/micro/go-micro"
	pb "github.com/naichadouban/gomicroservices/consignment-service/proto/consignment"
	vesselProto "github.com/naichadouban/gomicroservices/vessel-service/proto/vessel"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	// 获取容器设置的数据库地址环境变量
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = defaultHost
	}
	session, err := CreateSession(dbHost)
	if err != nil {
		log.Fatalf("create session error: %v\n", err)
	}
	defer session.Close()
	// set-up micro instance
	server := micro.NewService(
		micro.Name("shippy.service.consignment"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
	)
	// Init will parse the command line flags
	server.Init()
	// vessel-service 的客户端
	vesselClient := vesselProto.NewVesselServiceClient("shippy.service.vessel", server.Client())

	// register handler
	pb.RegisterShippingServiceHandler(server.Server(), &handler{session, vesselClient})
	// Run the server
	if err := server.Run(); err != nil {
		log.Panicf("micro server run error:%v\n", err)
	}
}

// AuthWrapper 是一个高阶函数，入参是 ”下一步“ 函数，出参是认证函数
// 在返回的函数内部处理完认证逻辑后，再手动调用 fn() 进行下一步处理
// token 是从 consignment-ci 上下文中取出的，再调用 user-service 将其做验证
// 认证通过则 fn() 继续执行，否则报错
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		llog.Tracef("AuthWrapper handler:%v", req)
		// consignment-service 独立测试时不进行认证
		if os.Getenv("DISABLE_AUTH") == "true" {
			return fn(ctx, req, resp)
		}
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			llog.Error("no auth meta-data found in request")
			return errors.New("no auth meta-data found in request")
		}
		llog.Tracef("AuthWrapper:Get metadata from context:%v", meta)
		// TODO
		// 注意Token字段必须大写（有知道的小伙伴请告知）
		token := meta["Token"]
		// 认证
		userClient := userProto.NewUserServiceClient("shippy.service.user", client.DefaultClient)
		authRes, err := userClient.ValidateToken(context.Background(), &userProto.Token{Token: token})
		if err != nil || authRes.Valid != true {
			llog.Errorf("AuthWrapper:ValidateToken error:%v", err)
			return err
		}

		err = fn(ctx, req, resp)
		return err
	}
}
