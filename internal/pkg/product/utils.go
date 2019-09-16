package product

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"kiren-backend-go/internal/app"
	"kiren-backend-go/internal/pkg/helper"
)

// loadProduct reads product.json and returns list of products.
// The function fetches the data from `mocked/products.json`.
// Otherwise, it fetches the data from S3 storage.
func loadProduct() (Products, error) {
	products := Products{}

	// Read product data
	if app.CFG.App.Env == "local" {
		// If the service is runnning on local environment, load products data from a local file.
		productFile, err := ioutil.ReadFile("mocked/products.json")
		if err != nil {
			return products, err
		}
		if err = json.Unmarshal([]byte(productFile), &products); err != nil {
			return products, err
		}
	} else {
		// If the service is not runnning on local environment, fetches products data from S3.
		fileByte, err := helper.GetFileFromS3("products.json")
		if err != nil {
			return products, err
		}
		if err = json.Unmarshal(fileByte, &products); err != nil {
			return products, err
		}
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

// getBestSellerProduct selects best seller products using best_seller_flag field.
func (products Products) getBestSellerProduct() []Product {
	bestSellerProducts := []Product{}

	// Append best seller products to bestSellerProducts slice
	for i := range products {
		if products[i].BestSellerFlag {
			bestSellerProducts = append(bestSellerProducts, products[i])
		}
	}

	return bestSellerProducts
}
