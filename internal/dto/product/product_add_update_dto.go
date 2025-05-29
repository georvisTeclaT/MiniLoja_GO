package product

type ProductAddUpdateDto struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
