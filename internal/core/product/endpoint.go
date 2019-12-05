package product

import (
	"database/sql"
	"kiren-backend-go/internal/app"
	"net/http"
)

// GetProducts is a logic of the /products endpoint.
// The function loads the product data and returns the value in the JSON form.
func GetProducts() (GetProductsResponse, int) {
	// Initialize logger
	logger := app.InitLogger()

	// Initialize response model
	res := GetProductsResponse{}

	selDB, err := app.Query(`
		SELECT product_id, product_name_th, product_description_th, product_image_path, product_price 
		FROM products 
	`)
	if err != nil {
		logger.Errorf("%s: %+v", "Query failed", err)
		res.Error = app.ErrorResp{
			Name:    app.ERR.InternalServerError.Name,
			Message: app.ERR.InternalServerError.Message,
		}
		return res, app.ERR.InternalServerError.Code
	}

	products := Products{}
	for selDB.Next() {
		var productID, productNameTH, productNameEN, productDescriptionTH, productDescriptionEN, productImagePath sql.NullString
		var productPrice sql.NullFloat64
		var productBestSellerFlag sql.NullBool
		err = selDB.Scan(&productID, &productNameTH, &productDescriptionTH, &productImagePath, &productPrice)
		if err != nil {
			logger.Errorf("%s: %+v", "Cannot fetch products data", err)
			res.Error = app.ErrorResp{
				Name:    app.ERR.InternalServerError.Name,
				Message: app.ERR.InternalServerError.Message,
			}
			return res, app.ERR.InternalServerError.Code
		}
		productSQLNull := ProductSQLNull{
			ID:               productID,
			NameTH:           productNameTH,
			NameEN:           productNameEN,
			DescriptionTH:    productDescriptionTH,
			DescriptionEN:    productDescriptionEN,
			ProductImagePath: productImagePath,
			Price:            productPrice,
			BestSellerFlag:   productBestSellerFlag,
		}
		product := MapSQLNullToProduct(productSQLNull)
		products = append(products, product)
	}

	logger.Debugf("Product data: %+v", products)
	res.ProductData = products
	return res, http.StatusOK
}

// GetProduct is a logic of the /product endpoint.
// The function inquiries a product and returns in JSON form.
func GetProduct(productID string) (GetProductResponse, int) {
	// Initialize logger
	logger := app.InitLogger()

	// Initialize response model
	res := GetProductResponse{}

	product := Product{}
	sqlRow, err := app.QueryRow(`
		SELECT product_id, product_name_th, product_description_th, product_image_path, product_price 
		FROM products 
		WHERE product_id="` + productID + `"`,
	)
	if err != nil {
		logger.Errorf("%s: %+v", "Cannot fetch product data", err)
		res.Error = app.ErrorResp{
			Name:    app.ERR.InternalServerError.Name,
			Message: app.ERR.InternalServerError.Message,
		}
		return res, app.ERR.InternalServerError.Code
	}
	sqlRow.Scan(&product.ID, &product.NameTH, &product.DescriptionTH, &product.ProductImagePath, &product.Price)

	logger.Debugf("Product data: %+v", product)
	res.ProductData = product

	return res, http.StatusOK
}

// GetBestSellerProducts is a logic of the /best_seller_product endpoint.
// The function inquiries best seller products and returns in JSON form.
func GetBestSellerProducts() (BestSellerProductResponse, int) {
	// Initialize logger
	logger := app.InitLogger()

	// Initialize response model
	res := BestSellerProductResponse{}

	// Load all products
	products, err := loadProduct()
	if err != nil {
		logger.Errorf("%s: %+v", "Cannot fetch product data", err)
		res.Error = app.ErrorResp{
			Name:    app.ERR.InternalServerError.Name,
			Message: app.ERR.InternalServerError.Message,
		}
		return res, app.ERR.InternalServerError.Code
	}

	// Filter best seller products using best_seller_flag
	bestSellerProducts := products.getBestSellerProduct()
	if err != nil {
		logger.Errorf("%s: %+v", "Cannot select product data", err)
		res.Error = app.ErrorResp{
			Name:    app.ERR.InternalServerError.Name,
			Message: app.ERR.InternalServerError.Message,
		}
		return res, app.ERR.InternalServerError.Code
	}

	res.ProductData = bestSellerProducts
	return res, http.StatusOK
}
