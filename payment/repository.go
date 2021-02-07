package payment

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

// Repository struct
type Repository struct {
	database *sql.DB
}

// NewRepository Create a new repository
func NewRepository(database *sql.DB) *Repository {
	return &Repository{database: database}
}

// Insert a new record
func (repository *Repository) Insert(payment Payment) (string, error) {
	tx, err := repository.database.Begin()
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
	_, err = stmt.Exec(id, payment.Amount, payment.Currency, payment.Source)
	if err != nil {
		log.Println(err)
		return "", err
	}
	tx.Commit()

	return id.String(), nil
}

// GetByID get a single payment
func (repository *Repository) GetByID (id string) (*Payment, error) {
	stmt, err := repository.database.Prepare("select * from payments where id = ?")
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

	return &Payment{
		ID: ID,
		Amount: amount,
		Currency: currency,
		Source: source,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

// Search payments
func (repository *Repository) Search() (*[]Payment, error) {
	payments := []Payment{}
	rows, err := repository.database.Query("select * from payments")
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

		payments = append(payments, Payment{
			ID: ID,
			Amount: amount,
			Currency: currency,
			Source: source,
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
