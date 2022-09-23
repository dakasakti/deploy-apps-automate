package user

import (
	"context"
	"errors"

	"github.com/dakasakti/deploy-apps-hexagonal/internal/factory"
	"github.com/dakasakti/deploy-apps-hexagonal/internal/models"
	"github.com/dakasakti/deploy-apps-hexagonal/internal/repository/user"
)

type UserService interface {
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUser(ctx context.Context, user_id uint) (*models.User, error)
	CreateUser(ctx context.Context, req models.UserRequest) (interface{}, error)
	UpdateUser(ctx context.Context, user_id uint, req models.UserUpdateRequest) error
	DeleteUser(ctx context.Context, user_id uint) error
	Login(ctx context.Context, req models.UserLogin) (*string, error)
}

type userService struct {
	ur user.UserRepository
}

func NewUserService(f *factory.Factory) UserService {
	return &userService{
		ur: f.UserRepository,
	}
}

func (us *userService) GetUsers(ctx context.Context) ([]models.User, error) {
	results, err := us.ur.Gets(ctx)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, errors.New("data tidak ditemukan")
	}

	return results, nil
}

func (us *userService) GetUser(ctx context.Context, user_id uint) (*models.User, error) {
	result, err := us.ur.Get(ctx, user_id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *userService) CreateUser(ctx context.Context, req models.UserRequest) (interface{}, error) {
	data := models.User{
		Fullname: req.Fullname,
		Username: req.Username,
		Password: req.Password,
	}

	result, err := us.ur.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *userService) UpdateUser(ctx context.Context, user_id uint, req models.UserUpdateRequest) error {
	if req.Fullname == "" && req.Username == "" && req.Password == "" {
		return errors.New("tidak ada yang diperbaiki")
	}

	data := models.User{
		ID:       user_id,
		Fullname: req.Fullname,
		Username: req.Username,
		Password: req.Password,
	}

	err := us.ur.Update(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (us *userService) DeleteUser(ctx context.Context, user_id uint) error {
	err := us.ur.Delete(ctx, user_id)
	if err != nil {
		return err
	}

	return nil
}

func (us *userService) Login(ctx context.Context, req models.UserLogin) (*string, error) {
	data := models.User{
		Username: req.Username,
		Password: req.Password,
	}

	token, err := us.ur.Login(data)
	if err != nil {
		return nil, err
	}

	return token, nil
}
