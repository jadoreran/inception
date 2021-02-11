package service_test

import (
	"log"
	"testing"

	"github.com/jadoreran/inception/domain"
	"github.com/jadoreran/inception/provider"
	"github.com/jadoreran/inception/repository"
)

// Provider struct
type omiseTest struct {
}

// CreateCharge use type as source
func (*omiseTest) CreateCharge(amount int64, currency string, sourceType string) error {
	return nil
}

// fakePaymentRepository struct
type paymentRepositoryTest struct {
}

// GetByID get a single payment
func (*paymentRepositoryTest) Insert(p *domain.Payment) (string, error) {
	return "", nil
}

// GetByID get a single payment
func (*paymentRepositoryTest) GetByID(id string) (*domain.Payment, error) {
	return nil, nil
}

// GetByID get a single payment
func (*paymentRepositoryTest) Search() (*[]domain.Payment, error) {
	return nil, nil
}

// fakePaymentServicer struct
type paymentServiceTest struct {
	repository repository.Store
	provider   provider.Provider
}

// newFakePaymentService Create a new repository
func newPaymentServiceTest(r repository.Store, p provider.Provider) *paymentServiceTest {
	return &paymentServiceTest{repository: r, provider: p}
}

// CreatePayment a new payment
func (s *paymentServiceTest) CreatePayment(p *domain.Payment) (string, error) {
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

func TestInsert(t *testing.T) {
	repository := &paymentRepositoryTest{}
	provider := &omiseTest{}
	service := newPaymentServiceTest(repository, provider)

	payment := domain.Payment{}

	_, err := service.CreatePayment(&payment)
	if err != nil {
		t.Error(err)
	}
}
