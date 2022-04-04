package entity

type Category struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
}