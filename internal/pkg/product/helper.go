package product

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"kiren-backend-go/internal/app"
	"kiren-backend-go/internal/pkg/helper"
)

// loadProduct reads product.json and returns list of products.
func loadProduct() (Products, error) {
	products := Products{}

	// Read product data
	if app.CFG.App.Env == "local" {
		// If the service is runnning on local environment, load products data from a local file.
		productFile, err := ioutil.ReadFile("data/products.json")
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
		fmt.Println(fileByte)
	}

	// Extract the JSON file and map it to the product model
	// err := json.Unmarshal([]byte(productFile), &products)
	// if err != nil {
	// 	return products, err
	// }

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
