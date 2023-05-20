package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		return nil, err
	}

	return db, nil
}




func TestCreateNewProduct(t *testing.T) {
	db, err := SetupTestDB()
	if err != nil {
		t.Fatal(err)
	}

	product , _:= entity.NewProduct("pc", 20.0)
	ProductDb := NewProduct(db)

	err = ProductDb.Create(product)
	assert.Nil(t, err )

	var productFound entity.Product
	err = db.First(&productFound, "id = ?", product.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, productFound.Name, productFound.Name)
	assert.Equal(t, productFound.Price, productFound.Price)
	assert.NotNil(t, productFound.Name)
	assert.NotNil(t, productFound.Price)

}


func TestFIndProductByID(t *testing.T) {
	db, err := SetupTestDB()
	if err != nil {
		t.Fatal(err)
	}

	product , _:= entity.NewProduct("pc", 20.0)
	ProductDb := NewProduct(db)

	err = ProductDb.Create(product)
	assert.Nil(t, err )

	productFound, err := ProductDb.FindByID(product.ID.String())
	assert.Nil(t, err )
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, productFound.Name, productFound.Name)
	assert.Equal(t, productFound.Price, productFound.Price)
	assert.NotNil(t, productFound.Name)
	assert.NotNil(t, productFound.Price)

}


func TestUpdateProduct(t *testing.T) {
	db, err := SetupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	product , _:= entity.NewProduct("pc", 20.0)
	ProductDb := NewProduct(db)
	err = ProductDb.Create(product)
	assert.Nil(t, err )
	err = ProductDb.Update(product)
	assert.NoError(t, err)
	product, err = ProductDb.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "pc", product.Name)
	
}

func TestFindByAllProducts(t *testing.T) {
	db, err := SetupTestDB()
	if err != nil {
		t.Fatal(err)
	}

	for i := 1; i < 24; i++ {
		Product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(Product)
}
	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)


	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)

}


func TestDeleteProduct(t *testing.T) {
	db, err := SetupTestDB()
	if err != nil {
		t.Fatal(err)
	}

    product, _:= entity.NewProduct("pc", 20.0)
    ProductDb := NewProduct(db)
    err = ProductDb.Create(product)
    assert.Nil(t, err )
    err = ProductDb.Delete(product.ID.String())
    assert.NoError(t, err)
    assert.Nil(t, err)


}