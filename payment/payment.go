package payment

import "time"

// Payment struct
type Payment struct {
	ID        string     `json:"id"`
	Amount    int        `json:"amount"`
	Currency  string     `json:"currency"`
	Source    string     `json:"source"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	CreatedAt *time.Time `db:"created_at" json:"createdAt"`
}
