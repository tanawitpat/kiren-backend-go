package product

import (
	"encoding/json"
	"io/ioutil"
)

// loadProduct reads product.json and returns list of products.
func loadProduct() ([]Product, error) {
	products := []Product{}

	// Read product data
	productFile, err := ioutil.ReadFile("data/products.json")
	if err != nil {
		return products, err
	}

	// Extract the JSON file and map it to the product model
	err = json.Unmarshal([]byte(productFile), &products)
	if err != nil {
		return products, err
	}

	return products, nil
}
