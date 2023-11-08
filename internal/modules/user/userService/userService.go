package userservice

import (
	userrepository "crud/internal/modules/user/userRepository"
)

type UserService struct {
	userRepo userrepository.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepo: userrepository.New(),
	}
}
