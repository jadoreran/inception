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
func (repository *Repository) Insert(db *sql.DB, name string) string {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(`INSERT INTO payments(id, amount, created_at, updated_at) VALUES (?, ?, datetime('now'), datetime('now'))`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	id := uuid.New()
	_, err = stmt.Exec(id, 500)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()

	return id.String()
}

// GetByID get a single payment
func (repository *Repository) GetByID (db *sql.DB, id string) Payment {
	stmt, err := db.Prepare("select * from payments where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var ID string
	var amount int
	var createdAt *time.Time
	var updatedAt *time.Time
	err = stmt.QueryRow(id).Scan(&ID, &amount, &createdAt, &updatedAt )
	if err != nil {
		log.Fatal(err)
	}

	return Payment{
		ID: ID,
		Amount: amount,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

// Search payments
func (repository *Repository) Search(db *sql.DB) *[]Payment{
	payments := []Payment{}
	rows, err := db.Query("select * from payments")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	for rows.Next() {
		var ID string
		var amount int
		var createdAt *time.Time
		var updatedAt *time.Time
		err = rows.Scan(&ID, &amount, &createdAt, &updatedAt )
		if err != nil {
			log.Fatal(err)
		}

		payments = append(payments, Payment{
			ID: ID,
			Amount: amount,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &payments
}
