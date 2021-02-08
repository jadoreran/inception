package payment

import (
	"log"

	"github.com/jadoreran/inception/provider"
)

// Service struct
type Service struct {
	repository *Repository
}

// NewPaymentService Create a new repository
func NewPaymentService(repository *Repository) *Service {
	return &Service{repository: repository}
}

// CreatePayment a new payment
func (service *Service) CreatePayment(payment Payment) (string, error) {
	p := provider.Omise{}
	err := p.CreateCharge(int64(payment.Amount), payment.Currency, payment.Source)
	if err != nil {
		log.Println(err)
		return "", err
	}

	id, err := service.repository.Insert(payment)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return id, nil
}

// FindPaymentByID find a single payment record
func (service *Service) FindPaymentByID(id string) (*Payment, error) {
	payment, err := service.repository.GetByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return payment, nil
}

// SearchPayments and return list of payments
func (service *Service) SearchPayments() (*[]Payment, error) {
	payments, err := service.repository.Search()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return payments, nil
}
