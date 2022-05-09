package users

import (
	"baby-daily-api/internal/model"
	"baby-daily-api/internal/repository/cache"
	"baby-daily-api/internal/repository/database/users"
)

type userService struct {
	userRepository users.UserRepository
	cache          cache.Cache
}

type UserService interface {
	Create(user *model.User) (*model.User, error)
	GetUserByOpenId(string) (*model.User, error)
	GetUserByUnionId(unionId string) (*model.User, error)
	Update(string, *model.User) (*model.User, error)
}

func NewUserService(userRepository users.UserRepository, cache cache.Cache) UserService {
	return &userService{
		userRepository: userRepository,
		cache:          cache,
	}
}

func (u *userService) Create(user *model.User) (*model.User, error) {
	return u.userRepository.Create(user)
}

func (u *userService) GetUserByOpenId(openId string) (*model.User, error) {
	return u.userRepository.GetUserByOpenId(openId)
}

func (u *userService) GetUserByUnionId(unionId string) (*model.User, error) {
	return u.userRepository.GetUserByUnionId(unionId)
}

func (u *userService) Update(openId string, user *model.User) (*model.User, error) {
	old, err := u.GetUserByOpenId(openId)
	if err != nil {
		return nil, err
	}
	user.ID = old.ID
	return u.userRepository.Update(user)
}
