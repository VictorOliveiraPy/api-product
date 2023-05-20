package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/VictorOliveiraPy/internal/dto"
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/VictorOliveiraPy/internal/infra/database"
)


type ProductHandler struct {
	ProductDB database.ProductInterface

}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
        ProductDB: db,
    }
}


func (handler *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request){
	var product dto.CreateProductInput

	err := json.NewDecoder(r.Body).Decode(&product)
	if err!= nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	newProduct, err := entity.NewProduct(product.Name, product.Price)
	if err!= nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	err = handler.ProductDB.Create(newProduct)
	if err!= nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }


	w.WriteHeader(http.StatusCreated)
}