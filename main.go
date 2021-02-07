package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

//Payment struct
type Payment struct {
	ID      string    `json:"id"`
	Amount       int       `json:"amount"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updatedAt"`
	CreatedAt   *time.Time `db:"created_at" json:"createdAt"`
}

func main() {
	os.Remove("inception.db") 
	file, err := os.Create("inception.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()

	db, _ := sql.Open("sqlite3", "./inception.db")
	defer db.Close()

	createTable(db)

	r := gin.Default()
	r.POST("/payment", func(c *gin.Context) {
		id := insert(db, "ssss")

	c.JSON(200, gin.H{
		"message": id,
	})
	})
	
	r.GET("/payment/:id", func(c *gin.Context) {
		id := c.Param("id")
		payment := queryByID(db, id)

		c.JSON(200, gin.H{
			"payload": payment,
		})
	})

	r.GET("/payments", func(c *gin.Context) {
		payments := find(db)

		c.JSON(200, gin.H{
			"payload": payments,
		})
	})
	r.Run()
}

func insert(db *sql.DB, name string) string {
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

func queryByID (db *sql.DB, id string) Payment {
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

func find(db *sql.DB) *[]Payment{
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

func createTable(db *sql.DB) {
sqlStmt := `
	create table payments (
		id text PRIMARY KEY, 
		amount integer not null, 
		created_at datetime not null, 
		updated_at datetime not null);
	`

	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}