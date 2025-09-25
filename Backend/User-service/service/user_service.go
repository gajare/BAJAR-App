package service

import (
	"context"
	"time"
	"user-service/models"
	"user-service/repository"
	"user-service/utils"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	return s.repo.CreateUser(ctx, user)
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	return s.repo.UpdateUser(ctx, user)
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.repo.DeleteUser(ctx, id)
}

// ListUsers returns all users
func (s *UserService) ListUsers(ctx context.Context) ([]*models.User, error) {
	return s.repo.ListUsers(ctx)
}

// GetUserByEmail retrieves a user by email
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}

// AuthenticateUser checks user credentials
func (s *UserService) AuthenticateUser(ctx context.Context, email, password string) (*models.User, error) {
	return s.repo.GetUserByEmailAndPassword(ctx, email, password)
}

// ChangePassword changes a user's password
func (s *UserService) ChangePassword(ctx context.Context, id string, newPassword string) error {
	return s.repo.ChangePassword(ctx, id, newPassword)
}

// ResetPassword resets a user's password
func (s *UserService) ResetPassword(ctx context.Context, email, newPassword string) error {
	return s.repo.ResetPassword(ctx, email, newPassword)
}

// Register registers a new user
func (s *UserService) Register(ctx context.Context, req *models.UserRequest) (*models.UserResponse, error) {
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hash,
		Phone:    req.Phone,
		Address:  req.Address,
	}
	if err := s.repo.CreateUser(ctx, &user); err != nil {
		return nil, err
	}
	resp := &models.UserResponse{ID: user.ID, Email: user.Email}
	return resp, nil
}

// Login logs in a user
func (s *UserService) Login(ctx context.Context, req *models.UserRequest) (*models.UserResponse, error) {
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if err := utils.CheckPassword(user.Password, req.Password); err != nil {
		return nil, err
	}
	token, _ := utils.CreateToken(user.ID, 24*time.Hour)
	resp := &models.UserResponse{ID: user.ID, Email: user.Email, Token: token}
	return resp, nil
}
