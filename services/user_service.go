package services

import (
	"volt/app/security"
	"volt/data/entitys"
	"volt/data/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository:userRepository}
}

func (this *UserService) GetByID(id int32) (*entitys.User, error) {
	return this.userRepository.GetById(id)
}

func (this *UserService) Login(username,pwd string)(*entitys.User,bool,string) {
	user,err := this.userRepository.GetByName(username)
	if err!=nil	{
		return nil,false,"用户名或密码错误"
	}
	sign,_ :=security.DesEncrypt(pwd)
	if user.PassWord!=sign {
		return nil,false,"用户名或密码错误"
	}
    return user,true,""
}