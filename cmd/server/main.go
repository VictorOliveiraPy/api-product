package main

import (
	"log"
	"net/http"

	"github.com/VictorOliveiraPy/configs"
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/VictorOliveiraPy/internal/infra/database"
	"github.com/VictorOliveiraPy/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err!= nil {
        panic(err)
    }

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err!= nil {
        panic(err)
    }

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	userDB := database.NewUser(db)

	ProductHandler := handlers.NewProductHandler(productDB)
	UserHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JwtExperesIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", ProductHandler.CreateProduct)
	r.Get("/products/{id}", ProductHandler.GetProduct)
	r.Put("/products/{id}", ProductHandler.UpdateProduct)
	r.Delete("/products/{id}", ProductHandler.DeleteProduct)
	r.Get("/products", ProductHandler.GetProducts)


	r.Post("/users",UserHandler.Create)
	r.Post("/users/generate_token", UserHandler.GetJWTUser)
	
	http.ListenAndServe(":8000", r)
}

