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
	logger.Debugf("Product data: %+v", products)

	// Select a product using product ID
	product, err := products.selectProduct(productID)
	if err != nil {
		logger.Errorf("%s: %+v", "Cannot select product data", err)
		res.Error = app.ErrorResp{
			Name:    app.ERR.InternalServerError.Name,
			Message: app.ERR.InternalServerError.Message,
		}
		c.JSON(app.ERR.InternalServerError.Code, res)
		return
	}

	relevantProducts, err := products.selectProductRandom(4)
	if err != nil {
		logger.Errorf("%s: %+v", "Cannot sample products from product data", err)
		res.Error = app.ErrorResp{
			Name:    app.ERR.InternalServerError.Name,
			Message: app.ERR.InternalServerError.Message,
		}
		c.JSON(app.ERR.InternalServerError.Code, res)
		return
	}

	res.ProductData = product
	res.RelevantProductData = relevantProducts

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
