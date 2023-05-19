package database

import "github.com/VictorOliveiraPy/internal/entity"


type UserInterface interface {
	Create(user *entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)

}