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

	// Query products data
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

	// Map the queried data with the products model
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
	res.ProductData = products
	logger.Debugf("Product data: %+v", products)

	return res, http.StatusOK
}

// GetProduct is a logic of the /product endpoint.
// The function inquiries a product and returns in JSON form.
func GetProduct(productID string) (GetProductResponse, int) {
	// Initialize logger
	logger := app.InitLogger()

	// Initialize response models
	res := GetProductResponse{}
	product := Product{}
	relevantProducts := Products{}

	// Return error if no product ID provided
	if productID == "" {
		logger.Errorf("%s", "No productID provided")
		res.Error = app.ErrorResp{
			Name:    app.ERR.BadRequest.Name,
			Message: app.ERR.BadRequest.Message,
		}
		return res, app.ERR.BadRequest.Code
	}

	// Query product data using product ID and map the data with the product model
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
	res.ProductData = product

	// Select products randomly for product recommendation
	sqlRowRelevantProducts, err := app.Query(`
		SELECT product_id, product_name_th, product_description_th, product_image_path, product_price
		FROM products
		ORDER BY RAND()
		LIMIT 4
	`)
	if err != nil {
		logger.Errorf("%s: %+v", "Cannot fetch relevant product data", err)
		res.Error = app.ErrorResp{
			Name:    app.ERR.InternalServerError.Name,
			Message: app.ERR.InternalServerError.Message,
		}
		return res, app.ERR.InternalServerError.Code
	}

	// Map the queried data with the products struct
	for sqlRowRelevantProducts.Next() {
		var productID, productNameTH, productNameEN, productDescriptionTH, productDescriptionEN, productImagePath sql.NullString
		var productPrice sql.NullFloat64
		var productBestSellerFlag sql.NullBool
		err = sqlRowRelevantProducts.Scan(&productID, &productNameTH, &productDescriptionTH, &productImagePath, &productPrice)
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
		relevantProduct := MapSQLNullToProduct(productSQLNull)
		relevantProducts = append(relevantProducts, relevantProduct)
	}
	res.RelevantProductData = relevantProducts
	logger.Debugf("Relavant product data: %+v", relevantProducts)

	return res, http.StatusOK
}

// GetBestSellerProducts is a logic of the /best_seller_product endpoint.
// The function inquiries best seller products and returns in JSON form.
func GetBestSellerProducts() (BestSellerProductResponse, int) {
	// Initialize logger
	logger := app.InitLogger()

	// Initialize response model
	res := BestSellerProductResponse{}

	// Query products data
	selDB, err := app.Query(`
		SELECT product_id, product_name_th, product_description_th, product_image_path, product_price 
		FROM products
		WHERE best_seller_flag="Y"
	`)
	if err != nil {
		logger.Errorf("%s: %+v", "Query failed", err)
		res.Error = app.ErrorResp{
			Name:    app.ERR.InternalServerError.Name,
			Message: app.ERR.InternalServerError.Message,
		}
		return res, app.ERR.InternalServerError.Code
	}

	// Map the queried data with the products model
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
	res.ProductData = products
	logger.Debugf("Best seller product data: %+v", products)

	return res, http.StatusOK
}
