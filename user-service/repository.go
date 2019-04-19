package main

import (
	"github.com/jinzhu/gorm"
	pb "github.com/naichadouban/gomicroservices/user-service/proto/user"
	_ "github.com/lib/pq"
)

/**
repository：实现与数据库交互
 */
type repository interface {
	Create(*pb.User)error
	Get(id string)(*pb.User,error)
	GetAll()([]*pb.User,error)
	GetByEmailAndPassword(*pb.User) (*pb.User, error)
}
type UserRepository struct{
	db *gorm.DB
}
func (repo *UserRepository) Create(u *pb.User)error{
	return repo.db.Create(&u).Error
}
func (repo *UserRepository) Get(id string)(*pb.User,error){
	var u *pb.User
	if err := repo.db.Where("id = ?",id).First(&u).Error ;err != nil{
		return nil,err
	}
	return u,nil
}
func (repo *UserRepository) GetAll()([]*pb.User,error){
	var users []*pb.User
	if err := repo.db.Find(&users).Error;err != nil{
		return nil,err
	}
	return users,nil
}

func (repo *UserRepository) GetByEmailAndPassword(u *pb.User)(*pb.User,error){
	if err := repo.db.Find(&u).Error;err != nil{
		return nil,err
	}
	return u,nil
}