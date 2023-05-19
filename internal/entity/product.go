package entity

import (
	"errors"
	"time"

	"github.com/VictorOliveiraPy/pkg/entity"
)

var (
	ErrIDIsRequired = 		errors.New("id is required")
	ErrInvalidId	= 		errors.New("invalid id")
	ErrNameIsRequired = 	errors.New("name is required")
	ErrInvalidPrice = 		errors.New("invalid price")
	ErrPriceIsRequired = 	errors.New("price is required")
)



type Product struct {
	ID entity.ID `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	CreatedAt time.Time `json:"created_at"`

}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID: entity.NewId(),
		Name: name,
        Price: price,
        CreatedAt: time.Now(),
    }
	err := product.ValidateEntity()
	if err!= nil {
        return nil, err
    }
	return product, nil
	}


func (p *Product) ValidateEntity() error {
	if p.ID.String() == "" {
			return ErrIDIsRequired
		}

		if _, err := entity.ParseId(p.ID.String()); err != nil {
			return ErrInvalidId
		}
		if p.Name == "" {
			return ErrNameIsRequired
		}
		if p.Price == 0 {
			return ErrInvalidPrice
		}
		if p.Price <= 0 {
			return ErrInvalidPrice
        }
		return nil

}