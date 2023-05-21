package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/VictorOliveiraPy/internal/dto"
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/VictorOliveiraPy/internal/infra/database"
)


type UserHandler struct {
	userDB database.UserInterface

}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{
        userDB: userDB,
    }
}

func (handler *UserHandler) Create(w http.ResponseWriter, r *http.Request){
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err!= nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	u, err := entity.NewUser(user.Email, user.Name, user.Password)
	if err!= nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	err = handler.userDB.Create(u)
	if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	w.WriteHeader(http.StatusCreated)
	
}