package services

import (
	"github.com/mashingan/smapping"
	"github.com/octane77/rova/identityService/dtos"
	"github.com/octane77/rova/identityService/entities"
	"github.com/octane77/rova/identityService/repositories"
	"github.com/octane77/rova/identityService/utils"
)

type UserService interface {
	CreateUser(createUserDto dtos.CreateUserDto) (*entities.User, error)
	GetUser(id uint) (*entities.User, error)
	GetUserByName(name string) (*entities.User, error)
}

type userService struct {
	userRepository repositories.Repository[entities.User]
}

func (u userService) CreateUser(createUserDto dtos.CreateUserDto) (*entities.User, error) {
	var user entities.User
	err := smapping.FillStruct(&user, smapping.MapFields(&createUserDto))
	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPwd
	if err != nil {
		return nil, err
	}
	return u.userRepository.Create(&user)
}

func (u userService) GetUser(id uint) (*entities.User, error) {
	return u.userRepository.FindById(id, nil)
}

func (u userService) GetUserByName(name string) (*entities.User, error) {
	return u.userRepository.FindOne(&entities.User{Name: name}, nil)
}

func NewUserService(userRepository repositories.Repository[entities.User]) UserService {
	return userService{
		userRepository: userRepository,
	}
}
