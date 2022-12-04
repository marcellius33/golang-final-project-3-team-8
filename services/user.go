package services

import (
	"golang.org/x/crypto/bcrypt"
	"kanbanboard/helpers"
	"kanbanboard/models"
	"kanbanboard/params"
	"kanbanboard/repositories"
	"time"
)

type UserService interface {
	Register(userRegisterRequest params.UserRegisterRequest) (*params.UserRegisterResponse, error)
	Login(userLoginRequest params.UserLoginRequest) (*params.UserLoginResponse, error)
	Update(id uint, userUpdateRequest params.UserUpdateRequest) (*params.UserUpdateResponse, error)
	Delete(id uint) error
}

type userService struct {
	repository repositories.UserRepository
}

func (u userService) Register(userRegisterRequest params.UserRegisterRequest) (*params.UserRegisterResponse, error) {
	pwHash, err := bcrypt.GenerateFromPassword([]byte(userRegisterRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return &params.UserRegisterResponse{}, err
	}

	newUser := models.User{
		FullName: userRegisterRequest.FullName,
		Email:    userRegisterRequest.Email,
		Password: string(pwHash),
		Role:     "member",
	}

	_, err = u.repository.CreateUser(&newUser)

	if err != nil {
		return &params.UserRegisterResponse{}, err
	}
	resp := params.ParseToCreateUserResponse(&newUser)

	return &resp, nil
}

func (u userService) Login(userLoginRequest params.UserLoginRequest) (*params.UserLoginResponse, error) {
	userFound, err := u.repository.FindUserByEmail(userLoginRequest.Email)
	if err != nil {
		return &params.UserLoginResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(userLoginRequest.Password))
	if err != nil {
		return &params.UserLoginResponse{}, err
	}

	token := helpers.GenerateToken(userFound.ID, userFound.Email)

	resp := params.UserLoginResponse{}
	resp.Token = token

	return &resp, nil
}

func (u userService) Update(id uint, userUpdateRequest params.UserUpdateRequest) (*params.UserUpdateResponse, error) {
	userModel, err := u.repository.FindUserByID(id)
	if err != nil {
		return &params.UserUpdateResponse{}, err
	}

	userModel.Email = userUpdateRequest.Email
	userModel.FullName = userUpdateRequest.FullName
	userModel.UpdatedAt = time.Now()
	user, err := u.repository.UpdateUser(id, userModel)
	if err != nil {
		return &params.UserUpdateResponse{}, err
	}

	resp := params.ParseToUpdateUserResponse(user)
	return &resp, err
}

func (u userService) Delete(id uint) error {
	return u.repository.DeleteUser(id)
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}
