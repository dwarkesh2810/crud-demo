package userservice

import (
	userdto "crud/internal/modules/user/userDto"
	usermodel "crud/internal/modules/user/userModel"
	userrepository "crud/internal/modules/user/userRepository"
	userrequest "crud/internal/modules/user/userRequest"
	userresponse "crud/internal/modules/user/userResponse"
	"crud/pkg/helper"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo userrepository.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepo: userrepository.New(),
	}
}

func (userService *UserService) Create(request userrequest.RegisterRequest) (userresponse.UserResponse, error) {
	var response userresponse.UserResponse
	var user usermodel.Users

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)

	if err != nil {
		return response, errors.New("error hashing password")
	}
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email
	user.Password = string(hashPassword)
	user.CreatedAt = time.Now().Unix()
	user.UserId = helper.GetNewUserId()
	user.UserRole = "user"

	newUser, err := userService.userRepo.Create(user)

	if err != nil {
		return response, errors.New("error on creating user")
	}

	helper.SendMail(request.Email, os.Getenv("SUBJECT"), os.Getenv("MESSAGE"))
	return userdto.ToUser(newUser), nil
}

func (userService *UserService) CheckUserExist(email string) bool {
	user := userService.userRepo.FindByEmail(email)
	return user.ID != 0
}

func (userService *UserService) Login(request userrequest.LoginRequest) (userresponse.LoginUserResponse, error) {
	var response userresponse.LoginUserResponse
	existUser := userService.userRepo.FindByEmail(request.Email)

	if existUser.ID == 0 {
		return response, errors.New("invalid credentials")
	}

	ok, _ := helper.VerifyHashedData(request.Password, existUser.Password)

	if !ok {
		return response, errors.New("invalid credentials")
	}

	// Generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": existUser.UserId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	accessToken, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return response, errors.New("error while generating token")
	}

	return userdto.ToUserLogin(existUser, accessToken), nil

}
