package database

import "github.com/VictorOliveiraPy/internal/entity"


type UserInterface interface {
	Create(user *entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)

}


type ProductInterface interface {
	Create(product *entity.Product) (*entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Update(product *entity.Product) (*entity.Product, error)
	Delete(id string) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
}