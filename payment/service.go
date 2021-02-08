package payment

import (
	"log"
	"time"

	"github.com/jadoreran/inception/provider"
)

// Payment struct
type Payment struct {
	ID        string     `json:"id"`
	Amount    int        `json:"amount"`
	Currency  string     `json:"currency"`
	Source    string     `json:"source"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	CreatedAt *time.Time `db:"created_at" json:"createdAt"`
}

// New Create a new payment object
func New(amount int, currency string, source string) *Payment {
	return &Payment{
		Amount:   amount,
		Currency: currency,
		Source:   source,
	}
}

// Service struct
type Service struct {
	repository *Repository
}

// NewService Create a new repository
func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

// CreatePayment a new payment
func (service *Service) CreatePayment(payment *Payment) (string, error) {
	p := provider.New()
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
