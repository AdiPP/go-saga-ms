package entity

import "github.com/guregu/null"

type Product struct {
	ID		     int      `gorm:"primaryKey"`
	Title      string   
	Price      float64  
	CategoryID null.Int 			
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type JsonProduct struct {
	ID		     int      `json:"id"`
	Title      string   `json:"title"`
	Price      float64  `json:"price"`
	CategoryID null.Int `json:"category_id"`
	Category 	 Category `json:"category"`
}

func MapProductToJson(product *Product) JsonProduct {
	return JsonProduct {
		ID: product.ID,
		Title: product.Title,
		Price: product.Price,
		CategoryID: product.CategoryID,
		Category: product.Category,
	}
}