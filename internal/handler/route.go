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
		pingpong.Ping(c)
	})

	r.GET("/products", func(c *gin.Context) {
		product.GetProducts(c)
	})

	r.GET("/product/:id", func(c *gin.Context) {
		product.GetProduct(c, c.Param("id"))
	})

	r.GET("/best-seller-products", func(c *gin.Context) {
		product.GetBestSellerProducts(c)
	})

	return r
}
