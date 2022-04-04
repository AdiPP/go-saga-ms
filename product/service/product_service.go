package service

import (
	"errors"

	"github.com/AdiPP/e-commerce/product/entity"
	"github.com/AdiPP/e-commerce/product/repository"
)

type ProductService interface {
	Validate(product *entity.Product) error
	Create(product *entity.Product) (*entity.Product, error)
	FindAll()  ([]entity.Product, error)
}

var (
	productRepository repository.Repository = repository.NewProductRepository()
)

type Service struct {}

func NewProductService() ProductService {
	return &Service{}
}

func (s *Service) Validate(product *entity.Product) error {
	if product == nil {
		err := errors.New("the product is empty")
		return err
	}

	if product.Title == "" {
		err := errors.New("the product title is empty")
		return err
	}

	if product.Price == 0 {
		err := errors.New("the product price can't be zero")
		return err
	}

	return nil
}

func (s *Service) Create(product *entity.Product) (*entity.Product, error) {
	product, err := productRepository.Save(product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Service) FindAll()  ([]entity.Product, error) {
	products, err := productRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return products, nil	
}