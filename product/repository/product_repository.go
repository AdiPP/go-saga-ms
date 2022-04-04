package repository

import (
	"log"
	"os"

	"github.com/AdiPP/e-commerce/product/database"
	"github.com/AdiPP/e-commerce/product/entity"
)

type ProductRepository struct {}

var (
	db database.Database = database.NewPostgresSqlDatabase()
)

func NewProductRepository() Repository {
	return &ProductRepository{}
}

func (r *ProductRepository) Save(post *entity.Product) (*entity.Product, error) {
	db, connErr := db.GetDBConnection()

	if connErr != nil {
		return nil, connErr
	}

	result := db.Create(&post)

	if result.Error != nil {
		return nil, result.Error
	}

	return post, nil
}

func (r *ProductRepository) FindAll() ([]entity.Product, error) {
	db, connErr := db.GetDBConnection()
	
	if connErr != nil {
		return nil, connErr
	}

	products := []entity.Product{}

	result := db.Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func check(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}