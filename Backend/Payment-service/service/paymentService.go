package service

import (
	"errors"
	"payment-service/models"
	"payment-service/repository"
)

type PaymentService interface {
	CreatePayment(models.Payment) (models.Payment, error)
	GetPaymentByID(uint64) (models.Payment, error)
	ListPayments() ([]models.Payment, error)
	UpdatePaymentStatus(uint64, string) (models.Payment, error)
	RefundPayment(uint64) (models.Payment, error)
}

type paymentService struct {
	repo repository.PaymentRepository
}

func NewPaymentService(r repository.PaymentRepository) PaymentService {
	return &paymentService{repo: r}
}

func (s *paymentService) CreatePayment(p models.Payment) (models.Payment, error) {
	p.Status = "pending"
	return s.repo.Create(p)
}

func (s *paymentService) GetPaymentByID(id uint64) (models.Payment, error) {
	return s.repo.GetByID(id)
}

func (s *paymentService) ListPayments() ([]models.Payment, error) {
	return s.repo.List()
}

func (s *paymentService) UpdatePaymentStatus(id uint64, status string) (models.Payment, error) {
	payment, err := s.repo.GetByID(id)
	if err != nil {
		return payment, err
	}
	payment.Status = status
	return s.repo.Update(payment)
}

func (s *paymentService) RefundPayment(id uint64) (models.Payment, error) {
	payment, err := s.repo.GetByID(id)
	if err != nil {
		return payment, err
	}
	if payment.Status != "success" {
		return payment, errors.New("only successful payments can be refunded")
	}
	payment.Status = "refunded"
	return s.repo.Update(payment)
}
