package main

import (
	"context"
	"log"
)
import  pb "github.com/naichadouban/gomicroservices/user-service/proto/user"
type handler struct {
	repo repository
}
func (h *handler)Create(ctx context.Context,req *pb.User,res *pb.Response)error{
	log.Printf("user-service receive create request: %v\n",req)
	if err := h.repo.Create(req);err != nil{
		return err
	}
	res.User = req
	return nil
}

func (h *handler) Get(ctx context.Context,req *pb.User,res *pb.Response)error{
	log.Printf("user-service receive get request:%v\n",req)
	u,err := h.repo.Get(req.Id)
	if err != nil{
		return err
	}
	res.User = u
	return nil
}
func (h *handler) GetAll(ctx context.Context,req *pb.Request,res *pb.Response)error{
	users,err:= h.repo.GetAll()
	if err != nil{
		return err
	}
	res.Users = users
	return nil
}
func (h *handler) Auth(ctx context.Context,req *pb.User,res *pb.Token)error{
	_,err := h.repo.GetByEmailAndPassword(req)
	if err != nil{
		return err
	}
	res.Token = "`x_2nam"
	return nil
}
func (h *handler) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}