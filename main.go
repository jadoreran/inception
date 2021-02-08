package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jadoreran/inception/domain"
	"github.com/jadoreran/inception/repository"
	"github.com/jadoreran/inception/service"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	os.Remove("inception.db")
	file, err := os.Create("inception.db")
	if err != nil {
		log.Println(err.Error())
	}
	file.Close()

	db, _ := sql.Open("sqlite3", "./inception.db")
	defer db.Close()

	createTable(db)

	repository := repository.NewRepository(db)
	service := service.NewService(repository)

	r := gin.Default()
	r.POST("/payment", func(c *gin.Context) {
		data := &domain.Payment{}
		c.Bind(data)

		payment := domain.NewPayment(data.Amount, data.Currency, data.Source)
		id, err := service.CreatePayment(payment)
		if err != nil {
			log.Println(err)
			c.JSON(404, gin.H{
				"error": err,
			})
		} else {
			c.JSON(200, gin.H{
				"payload": id,
			})
		}
	})

	r.GET("/payment/:id", func(c *gin.Context) {
		id := c.Param("id")
		payment, err := service.FindPaymentByID(id)
		if err != nil {
			log.Println(err)
			c.JSON(404, gin.H{
				"error": err,
			})
		} else {
			c.JSON(200, gin.H{
				"payload": payment,
			})
		}
	})

	r.GET("/payments", func(c *gin.Context) {
		payments, err := service.SearchPayments()
		if err != nil {
			log.Println(err)
			c.JSON(404, gin.H{
				"error": err,
			})
		} else {
			c.JSON(200, gin.H{
				"payload": payments,
			})
		}
	})

	r.Run()
}

func createTable(db *sql.DB) {
	sqlStmt := `
		create table payments (
			id text PRIMARY KEY, 
			amount integer not null, 
			currency text not null, 
			source text not null, 
			created_at datetime not null, 
			updated_at datetime not null);
		`

	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Println(err.Error())
	}
	stmt.Exec()
}
