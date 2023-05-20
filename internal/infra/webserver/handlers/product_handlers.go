package handlers

import (
	"encoding/json"

	"net/http"

	"github.com/VictorOliveiraPy/internal/dto"
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/VictorOliveiraPy/internal/infra/database"
	entityPkg "github.com/VictorOliveiraPy/pkg/entity"

	"github.com/go-chi/chi/v5"
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

func (handler *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := handler.ProductDB.FindByID(id)
	if err!= nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (handler *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
    if id == "" {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    var product entity.Product
    err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product.ID, err = entityPkg.ParseId(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = handler.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
    err = handler.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

    w.WriteHeader(http.StatusOK)

}