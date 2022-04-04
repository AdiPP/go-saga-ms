package controller

import (
	"encoding/json"
	"net/http"

	"github.com/AdiPP/e-commerce/product/entity"
	"github.com/AdiPP/e-commerce/product/errors"
	"github.com/AdiPP/e-commerce/product/service"
)

var (
	productService service.ProductService = service.NewProductService()
)

type ProductController interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
}

type Controller struct {}

func NewProductController() ProductController {
	return &Controller{}
}

func (c *Controller) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	products, err := productService.FindAll()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)	
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error getting the Product."})
		return
	}	

	var resp = make([]entity.JsonProduct, len(products))

	for idx, product := range products {
		resp[idx] = entity.MapProductToJson(&product)
	}


	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *Controller) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	
	var product entity.Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if  err != nil {
		w.WriteHeader(http.StatusInternalServerError)	
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error unmarshaling request."})
		return
	}

	err = productService.Validate(&product)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)	
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	result, err := productService.Create(&product)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)	
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error saving the post."})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}