package models

type Product struct {
	Id    int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
