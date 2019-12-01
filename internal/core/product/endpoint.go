package product

import (
	"kiren-backend-go/internal/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProducts is a logic of the /products endpoint.
// The function loads the product data and returns the value in the JSON form.
func GetProducts(c *gin.Context) {
	// Initialize logger
	logger := app.InitLogger()

	// Initialize response model
	res := GetProductsResponse{}

	// Load all products
	products, err := loadProduct()
	if err != nil {
		logger.Errorf("%s: %+v", "Cannot fetch products data", err)
		res.Error = app.ErrorResp{
			Name:    app.ERR.InternalServerError.Name,
			Message: app.ERR.InternalServerError.Message,
		}
		c.JSON(app.ERR.InternalServerError.Code, res)
		return
	}
	logger.Debugf("Product data: %+v", products)

	res.ProductData = products
	c.JSON(http.StatusOK, res)
}

// GetProduct is a logic of the /product endpoint.
// The function inquiries a product and returns in JSON form.
func GetProduct(c *gin.Context, productID string) {
	// Initialize logger
	logger := app.InitLogger()

	// Initialize response model
	res := GetProductResponse{}

	product := Product{}
	sqlRow, err := app.QueryRow(`SELECT product_id, product_name_th, product_description_th, product_image_path, product_price FROM products WHERE product_id="` + productID + `"`)
	if err != nil {
		logger.Errorln("HELLOWORLD")
		logger.Errorf("%s: %+v", "Cannot fetch product data", err)
		res.Error = app.ErrorResp{
			Name:    app.ERR.InternalServerError.Name,
			Message: app.ERR.InternalServerError.Message,
		}
		c.JSON(app.ERR.InternalServerError.Code, res)
		return
	}
	sqlRow.Scan(&product.ID, &product.NameTH, &product.DescriptionTH, &product.ProductImagePath, &product.Price)

	logger.Debugf("Product data: %+v", product)
	res.ProductData = product

	c.JSON(http.StatusOK, res)
}

// GetBestSellerProducts is a logic of the /best_seller_product endpoint.
// The function inquiries best seller products and returns in JSON form.
func GetBestSellerProducts(c *gin.Context) {
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
		c.JSON(app.ERR.InternalServerError.Code, res)
		return
	}

	// Filter best seller products using best_seller_flag
	bestSellerProducts := products.getBestSellerProduct()
	if err != nil {
		logger.Errorf("%s: %+v", "Cannot select product data", err)
		res.Error = app.ErrorResp{
			Name:    app.ERR.InternalServerError.Name,
			Message: app.ERR.InternalServerError.Message,
		}
		c.JSON(app.ERR.InternalServerError.Code, res)
		return
	}

	res.ProductData = bestSellerProducts
	c.JSON(http.StatusOK, res)
}
