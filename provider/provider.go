package provider

import (
	"log"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

const (
	omisePublicKey = "pkey_test_5msppk1y7gre58fgshk"
	omiseSecretKey = "skey_test_5msppk1y79ktkwocson"
)

// Provider interface
type Provider interface {
	CreateCharge(amount int64, currency string, sourceType string) error
}

// Omise struct
type Omise struct {
}

// NewOmise a new Omise payment object
func NewOmise() *Omise {
	return &Omise{}
}

// CreateCharge use type as source
func (*Omise) CreateCharge(amount int64, currency string, sourceType string) error {
	client, err := omise.NewClient(omisePublicKey, omiseSecretKey)
	if err != nil {
		log.Println(err)
		return err
	}

	// Creates a charge from the token
	source, createSource := &omise.Source{}, &operations.CreateSource{
		Amount:   amount,
		Currency: currency,
		Type:     sourceType,
	}
	if err := client.Do(source, createSource); err != nil {
		log.Println(err)
		return err
	}

	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:    amount,
		Currency:  currency,
		ReturnURI: "http://www.example.com",
		Source:    source.ID,
	}

	if err := client.Do(charge, createCharge); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
