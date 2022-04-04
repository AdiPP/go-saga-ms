package main

import (
	"encoding/json"
	"net/http"

	"github.com/AdiPP/e-commerce/product/controller"
	"github.com/AdiPP/e-commerce/product/router"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
	productController controller.ProductController = controller.NewProductController()
)

func main () {
	const port string = ":8000"

	httpRouter.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "pong",
		})
	})

	httpRouter.Get("/products", productController.GetProducts)

	httpRouter.Post("/products", productController.CreateProduct)

	httpRouter.Serve(port)
}