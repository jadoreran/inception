package domain

import (
	"time"
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

// NewPayment Create a new payment object
func NewPayment(amount int, currency string, source string) *Payment {
	return &Payment{
		Amount:   amount,
		Currency: currency,
		Source:   source,
	}
}
