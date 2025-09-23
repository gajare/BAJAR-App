package repository

import (
	"payment-service/models"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(models.Payment) (models.Payment, error)
	GetByID(uint64) (models.Payment, error)
	List() ([]models.Payment, error)
	Update(models.Payment) (models.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Create(p models.Payment) (models.Payment, error) {
	if err := r.db.Create(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

func (r *paymentRepository) GetByID(id uint64) (models.Payment, error) {
	var payment models.Payment
	err := r.db.First(&payment, id).Error
	return payment, err
}

func (r *paymentRepository) List() ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Find(&payments).Error
	return payments, err
}

func (r *paymentRepository) Update(p models.Payment) (models.Payment, error) {
	err := r.db.Save(&p).Error
	return p, err
}
