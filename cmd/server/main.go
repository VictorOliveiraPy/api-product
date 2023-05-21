package main

import (
	"net/http"

	"github.com/VictorOliveiraPy/configs"
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/VictorOliveiraPy/internal/infra/database"
	"github.com/VictorOliveiraPy/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/go-chi/chi/v5"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err!= nil {
        panic(err)
    }

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err!= nil {
        panic(err)
    }

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	ProductHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Post("/products", ProductHandler.CreateProduct)
	r.Get("/products/{id}", ProductHandler.GetProduct)
	r.Put("/products/{id}", ProductHandler.UpdateProduct)
	r.Delete("/products/{id}", ProductHandler.DeleteProduct)
	r.Get("/products", ProductHandler.GetProducts)
	
	http.ListenAndServe(":8000", r)
}

