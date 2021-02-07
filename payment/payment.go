package payment

import "time"

// Payment struct
type Payment struct {
	ID      			string    `json:"id"`
	Amount       int       `json:"amount"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updatedAt"`
	CreatedAt   *time.Time `db:"created_at" json:"createdAt"`
}
