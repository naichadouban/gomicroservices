package main

import (
	"encoding/json"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	_ "github.com/micro/go-plugins/broker/nats"
	userPb "github.com/naichadouban/gomicroservices/user-service/proto/user"
)

const topic = "user.created"

func main() {
	srv := micro.NewService(
		micro.Name("shippy.service.email"),
		micro.Version("latest"),
	)
	srv.Init()

	pubSub := srv.Server().Options().Broker
	if err := pubSub.Connect(); err != nil {
		llog.Panicf("connect to broker error:%v", err)
	}
	// 订阅消息
	_, err := pubSub.Subscribe(topic, func(pub broker.Publication) error {
		llog.Tracef("subscribe:%v",pub.Message())
		var user *userPb.User
		if err := json.Unmarshal(pub.Message().Body, &user); err != nil {
			llog.Errorf("Unmarshal pub.Message().Body error:%v", err)
			return err
		}
		llog.Infof("receive event:%s,user:%v", topic, user)
		return sendEmail(user)
	})
	if err != nil {
		llog.Errorf("subscribe topic error:%v", err)
	}
	if err := srv.Run(); err != nil {
		llog.Errorf("service run error:%v", err)
	}

}
func sendEmail(user *userPb.User) error {
	llog.Infof("send email to user:%s", user.Name)
	return nil
}
