package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jadoreran/inception/domain"
)

// Store interface
type Store interface {
	Insert(p *domain.Payment) (string, error)
	GetByID(id string) (*domain.Payment, error)
	Search() (*[]domain.Payment, error)
}

// PaymentRepository struct
type PaymentRepository struct {
	database *sql.DB
}

// NewPaymentRepository Create a new repository
func NewPaymentRepository(database *sql.DB) *PaymentRepository {
	return &PaymentRepository{database: database}
}

// Insert a new record
func (r *PaymentRepository) Insert(p *domain.Payment) (string, error) {
	tx, err := r.database.Begin()
	if err != nil {
		log.Println(err)
		return "", err
	}

	stmt, err := tx.Prepare(`INSERT INTO payments(id, amount, currency, source, created_at, updated_at) VALUES (?, ?, ?, ?, datetime('now'), datetime('now'))`)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	id := uuid.New()
	_, err = stmt.Exec(id, p.Amount, p.Currency, p.Source)
	if err != nil {
		log.Println(err)
		return "", err
	}
	tx.Commit()

	return id.String(), nil
}

// GetByID get a single payment
func (r *PaymentRepository) GetByID(id string) (*domain.Payment, error) {
	stmt, err := r.database.Prepare("select * from payments where id = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	var ID string
	var amount int
	var currency string
	var source string
	var createdAt *time.Time
	var updatedAt *time.Time
	err = stmt.QueryRow(id).Scan(&ID, &amount, &currency, &source, &createdAt, &updatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &domain.Payment{
		ID:        ID,
		Amount:    amount,
		Currency:  currency,
		Source:    source,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

// Search payments
func (r *PaymentRepository) Search() (*[]domain.Payment, error) {
	payments := []domain.Payment{}
	rows, err := r.database.Query("select * from payments")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ID string
		var amount int
		var currency string
		var source string
		var createdAt *time.Time
		var updatedAt *time.Time
		err = rows.Scan(&ID, &amount, &currency, &source, &createdAt, &updatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		payments = append(payments, domain.Payment{
			ID:        ID,
			Amount:    amount,
			Currency:  currency,
			Source:    source,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &payments, nil
}
