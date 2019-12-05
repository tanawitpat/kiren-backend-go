package product

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"kiren-backend-go/internal/app"
	"kiren-backend-go/internal/helper"
)

// loadProduct reads product.json and returns list of products.
// The function fetches the data from `mocked/products.json` on local environment.
// Otherwise, it fetches the data from S3 storage.
func loadProduct() (Products, error) {
	products := Products{}

	// Read product data
	if app.CFG.App.Env == "local" {
		// If the service is runnning on local environment, it loads products data from a local file.
		productFile, err := ioutil.ReadFile("mocked/products.json")
		if err != nil {
			return products, err
		}
		if err = json.Unmarshal([]byte(productFile), &products); err != nil {
			return products, err
		}
	} else {
		// If the service is not runnning on local environment, It fetches products data from S3.
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

// selectProductRandom selects products randomly.
func (products Products) selectProductRandom(nProduct int) ([]Product, error) {
	var outputProducts []Product
	if len(products) < nProduct {
		return outputProducts, errors.New("nProduct is more than a number of products")
	}

	randomIntegerSlice := helper.GenerateRandomInteger(nProduct)
	for i := range randomIntegerSlice {
		outputProducts = append(outputProducts, products[i])
	}

	return outputProducts, nil
}

func MapSQLNullToProduct(productSQLNull ProductSQLNull) Product {
	id := helper.ConvertNullStringToString(productSQLNull.ID)
	nameTH := helper.ConvertNullStringToString(productSQLNull.NameTH)
	nameEN := helper.ConvertNullStringToString(productSQLNull.NameEN)
	descriptionTH := helper.ConvertNullStringToString(productSQLNull.DescriptionTH)
	descriptionEN := helper.ConvertNullStringToString(productSQLNull.DescriptionEN)
	price := helper.ConvertNullFloat64ToFloat64(productSQLNull.Price)
	imagePath := helper.ConvertNullStringToString(productSQLNull.ProductImagePath)
	bestSellerFlag := helper.ConvertNullBoolToBool(productSQLNull.BestSellerFlag)
	product := Product{
		ID:               id,
		NameTH:           nameTH,
		NameEN:           nameEN,
		DescriptionTH:    descriptionTH,
		DescriptionEN:    descriptionEN,
		ProductImagePath: imagePath,
		Price:            price,
		BestSellerFlag:   bestSellerFlag,
	}
	return product
}
