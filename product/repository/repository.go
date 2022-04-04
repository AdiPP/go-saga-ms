package repository

import "github.com/AdiPP/e-commerce/product/entity"

type Repository interface {
	Save(post *entity.Product) (*entity.Product, error)
	FindAll() ([]entity.Product, error)
}