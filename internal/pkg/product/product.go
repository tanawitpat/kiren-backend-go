package product

import (
	"encoding/json"
	"io/ioutil"
	"kiren-backend-go/internal/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProducts is a logic of the /products endpoint.
// The function opens the data/products.json file and returns the data in the JSON form.
func GetProducts(c *gin.Context) {
	logger := app.InitLogger()

	productFile, err := ioutil.ReadFile("data/products.json")
	if err != nil {
		logger.Errorf("Cannot fetch products data: %+v", err)
	}

	products := []Product{}
	err = json.Unmarshal([]byte(productFile), &products)
	if err != nil {
		logger.Errorf("Cannot extract products data: %+v", err)
	}

	logger.Debugf("Product Data: %+v", products)
	c.JSON(http.StatusOK, products)
}
