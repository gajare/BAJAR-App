package repository

import (
	"context"
	"user-service/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, userModel *models.User) error {
	return r.db.WithContext(ctx).Create(&userModel).Error
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user *models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, userModel *models.User) error {
	return r.db.WithContext(ctx).Save(&userModel).Error
}

func (r *UserRepository) DeleteUser(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, id).Error
}

func (r *UserRepository) ListUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByEmailAndPassword(ctx context.Context, email, password string) (*models.User, error) {
	var user *models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	// Password check should be done in service layer for security
	return user, nil
}

func (r *UserRepository) ChangePassword(ctx context.Context, id string, newPassword string) error {
	return r.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Update("password", newPassword).Error
}

func (r *UserRepository) ResetPassword(ctx context.Context, email, newPassword string) error {
	return r.db.WithContext(ctx).Model(&models.User{}).Where("email = ?", email).Update("password", newPassword).Error
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var user *models.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
