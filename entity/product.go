package entity

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ListProductsResponse struct {
	Products []Product `json:"products"`
}

func GetProductCacheKey(id string) string {
	return "product:" + id
}
