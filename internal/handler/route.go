package handler

import (
	"kiren-backend-go/internal/pkg/pingpong"
	"kiren-backend-go/internal/pkg/product"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		pingpong.Ping(c)
	})

	r.GET("/products", func(c *gin.Context) {
		product.GetProducts(c)
	})

	r.GET("/product/:id", func(c *gin.Context) {
		product.GetProduct(c, c.Param("id"))
	})

	return r
}
