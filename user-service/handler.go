package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/micro/go-micro/broker"
	_ "github.com/micro/go-plugins/broker/nats"
	pb "github.com/naichadouban/gomicroservices/user-service/proto/user"
	"golang.org/x/crypto/bcrypt"
)

const topic = "user.created"

type handler struct {
	repo         repository
	tokenService Authable
	PubSub       broker.Broker
}

func (h *handler) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	llog.Tracef("user-service receive create request: %v\n", req)
	// hash处理用户密码
	hashedpwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		llog.Errorf("genetateFromPassword error:%v", err)
		return err
	}
	req.Password = string(hashedpwd)
	if err := h.repo.Create(req); err != nil {
		return err
	}
	res.User = req

	if err := h.publishEvent(req); err != nil {
		llog.Errorf("user service publishEvent error:%v", err)
		return err
	}
	return nil
}

func (h *handler) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	llog.Infof("user-service receive Get request:%v\n", req)
	u, err := h.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = u
	return nil
}
func (h *handler) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	llog.Tracef("user service receive GetAll request:%v", req)
	users, err := h.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}
func (h *handler) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	llog.Tracef("user service receive Auth request:%v", req)
	u, err := h.repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}
	// 进行密码验证
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		llog.Errorf("can't pass password check:%v", err)
		return err
	}
	tokenStr, err := h.tokenService.Encode(u)
	if err != nil {
		llog.Errorf("generate token string error:%v", err)
		return err
	}
	res.Token = tokenStr
	return nil
}
func (h *handler) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	llog.Tracef("user service receive ValidateToken:%v", req)
	// decode token
	claims, err := h.tokenService.Decode(req.Token)
	if err != nil {
		llog.Errorf("tokenString decode error:%v", err)
		return err
	}
	if claims.User.Id == "" {
		return errors.New("invalid user")
	}
	res.Valid = true
	return nil
}

// 发送消息通知
func (h *handler) publishEvent(user *pb.User) error {
	llog.Tracef("publish event:%s,data:%v",topic,user)
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}
	msg := &broker.Message{
		Header: map[string]string{
			"id": user.Id,
		},
		Body: body,
	}

	// 发布user.created消息
	if err := h.PubSub.Publish(topic, msg); err != nil {
		llog.Errorf("publish event error:%v", err)
	}
	return nil
}
