package service

import (
	"gofiber/repository"

	"github.com/gofiber/fiber/v2"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return userService{userRepository: repository}
}

func (s userService) NewUser(request RequestNewUser) (*UserResponse, error) {
	if request.Username == "" || request.Email == "" {
		return nil, fiber.ErrUnprocessableEntity
	}

	user := repository.User{
		Username: request.Username,
		Email:    request.Email,
	}

	result, err := s.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	userReqponse := UserResponse{
		ID:       result.Id,
		Username: result.Username,
	}

	return &userReqponse, nil
}
