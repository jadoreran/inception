package service

import (
	"log"

	"github.com/jadoreran/inception/domain"
	"github.com/jadoreran/inception/provider"
	"github.com/jadoreran/inception/repository"
)

// Service struct
type Service struct {
	repository *repository.PaymentRepository
}

// NewService Create a new repository
func NewService(r *repository.PaymentRepository) *Service {
	return &Service{repository: r}
}

// CreatePayment a new payment
func (s *Service) CreatePayment(p *domain.Payment) (string, error) {
	provider := provider.New()
	err := provider.CreateCharge(int64(p.Amount), p.Currency, p.Source)
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
func (s *Service) FindPaymentByID(id string) (*domain.Payment, error) {
	payment, err := s.repository.GetByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return payment, nil
}

// SearchPayments and return list of payments
func (s *Service) SearchPayments() (*[]domain.Payment, error) {
	payments, err := s.repository.Search()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return payments, nil
}
