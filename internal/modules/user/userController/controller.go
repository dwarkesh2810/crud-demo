package usercontroller

import userservice "crud/internal/modules/user/userService"

type Controllers struct {
	userService userservice.UserServiceInterface
}

func New() *Controllers {
	return &Controllers{
		userService: userservice.New(),
	}
}
