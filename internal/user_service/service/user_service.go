package service

import (
	usermodel "aegis_task/internal/user_service/models"
	"aegis_task/internal/user_service/repositories"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	CreateUser(ctx echo.Context, req usermodel.User) error
	FindUserByID(ctx echo.Context, id uint) (*usermodel.User, error)
	UpdateUser(ctx echo.Context, req usermodel.User) error
	DeleteUser(ctx echo.Context, id uint) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(ctx echo.Context, req usermodel.User) error {
	req.Id = uuid.New().String()
	return s.repo.Create(req)
}

func (s *userService) FindUserByID(ctx echo.Context, id uint) (*usermodel.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(ctx echo.Context, req usermodel.User) error {
	return s.repo.Update(req)
}

func (s *userService) DeleteUser(ctx echo.Context, id uint) error {
	return s.repo.Delete(id)
}
