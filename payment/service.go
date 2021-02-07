package payment

import (
	"database/sql"
)

// Service struct
type Service struct {
	repository *Repository
}

// NewPaymentService Create a new repository
func NewPaymentService(repository *Repository) *Service {
	return &Service{ repository: repository }
}

// CreatePayment a new payment
func (service *Service) CreatePayment(db *sql.DB, name string) string {
	id := service.repository.Insert(db, name)
	return id
}

// FindPaymentByID find a single payment record
func (service *Service) FindPaymentByID(db *sql.DB, id string) Payment {
	return service.repository.GetByID(db, id)
}

// SearchPayments and return list of payments
func (service *Service) SearchPayments(db *sql.DB) *[]Payment{
	return service.repository.Search(db)
}
