package service

import (
	"log"

	"github.com/jadoreran/inception/domain"
	"github.com/jadoreran/inception/provider"
	"github.com/jadoreran/inception/repository"
)

// PaymentServicer interface
type PaymentServicer interface {
	CreatePayment(p *domain.Payment) (string, error)
}

// PaymentService struct
type PaymentService struct {
	repository repository.Store
	provider   provider.Provider
}

// NewService Create a new repository
func NewService(r repository.Store, p provider.Provider) *PaymentService {
	return &PaymentService{repository: r, provider: p}
}

// CreatePayment a new payment
func (s *PaymentService) CreatePayment(p *domain.Payment) (string, error) {
	err := s.provider.CreateCharge(int64(p.Amount), p.Currency, p.Source)
	if err != nil {
		log.Println(err)
		return "", err
	}

	id, err := s.repository.Insert(p)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return id, nil
}

// FindPaymentByID find a single payment record
func (s *PaymentService) FindPaymentByID(id string) (*domain.Payment, error) {
	payment, err := s.repository.GetByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return payment, nil
}

// SearchPayments and return list of payments
func (s *PaymentService) SearchPayments() (*[]domain.Payment, error) {
	payments, err := s.repository.Search()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return payments, nil
}
