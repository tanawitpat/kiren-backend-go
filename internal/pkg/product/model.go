package product

// Product represent the product model
type Product struct {
	ID            string  `json:"product_id"`
	NameTH        string  `json:"product_name_th"`
	NameEN        string  `json:"product_name_en"`
	DescriptionTH string  `json:"product_description_th"`
	DescriptionEN string  `json:"product_description_en"`
	Price         float64 `json:"product_price"`
}
