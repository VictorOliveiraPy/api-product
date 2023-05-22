package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/VictorOliveiraPy/internal/dto"
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/VictorOliveiraPy/internal/infra/database"
	"github.com/go-chi/jwtauth"
)


type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	userDB database.UserInterface
	Jwt *jwtauth.JWTAuth
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{
        userDB: userDB,
    }
	
}

func (handler *UserHandler) GetJWTUser(w http.ResponseWriter, r *http.Request){
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	JwtExperesIn := r.Context().Value("JwtExperesIn").(int)

	var user dto.GetJWTInput

	err := json.NewDecoder(r.Body).Decode(&user)
	if err!= nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
        return
    }

	u, err := handler.userDB.FindByEmail(user.Email)
	if err!= nil {
		w.WriteHeader(http.StatusNotFound)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
		}

	if !u.ValidatePasswordUser(user.Password){
		w.WriteHeader(http.StatusUnauthorized)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(JwtExperesIn)).Unix() ,
	})

	accessToken := struct {
		AccessToken string `json:"access_token"`

	}{
		AccessToken: tokenString,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}


func (handler *UserHandler) Create(w http.ResponseWriter, r *http.Request){
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err!= nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
        return
    }

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err!= nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
        return
    }
	err = handler.userDB.Create(u)
	if err!= nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
        return
    }
	w.WriteHeader(http.StatusCreated)
	
}