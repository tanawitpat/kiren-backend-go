package product

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// loadProduct reads product.json and returns list of products.
func loadProduct() (Products, error) {
	products := Products{}

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

// selectProduct selects a product using product ID
func (products Products) selectProduct(expectedProductID string) (Product, error) {
	product := Product{}

	// Filter products via product ID
	for i := range products {
		if products[i].ID == expectedProductID {
			return products[i], nil
		}
	}

	return product, errors.New("Not found")
}
