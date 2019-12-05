package handler

import (
	"kiren-backend-go/internal/core/pingpong"
	"kiren-backend-go/internal/core/product"

	"github.com/gin-gonic/gin"
)

// NewRouter handles all routes in the service.
func NewRouter() *gin.Engine {
	r := gin.Default()

	// Insert CORS middleware definition before any routes
	r.Use(corsMiddleware)

	// Define routes
	r.GET("/ping", func(c *gin.Context) {
		response, httpStatus := pingpong.Ping()
		c.JSON(httpStatus, response)
	})

	r.GET("/products", func(c *gin.Context) {
		response, httpStatus := product.GetProducts()
		c.JSON(httpStatus, response)
	})

	r.GET("/product/:id", func(c *gin.Context) {
		response, httpStatus := product.GetProduct(c.Param("id"))
		c.JSON(httpStatus, response)
	})

	r.GET("/best-seller-products", func(c *gin.Context) {
		response, httpStatus := product.GetBestSellerProducts()
		c.JSON(httpStatus, response)
	})

	return r
}
