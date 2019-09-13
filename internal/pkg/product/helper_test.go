package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Happy test case of selectProduct function.
func TestFilterProductFound(t *testing.T) {
	products := Products{
		Product{
			ID:     "AA001",
			NameTH: "เตาย่าง หมายเลข 1",
			NameEN: "Thai stove No.1",
		},
		Product{
			ID:     "AA002",
			NameTH: "เตาย่าง หมายเลข 2",
			NameEN: "Thai stove No.2",
		},
		Product{
			ID:     "AA003",
			NameTH: "เตาย่าง หมายเลข 3",
			NameEN: "Thai stove No.3",
		},
	}
	product, err := products.selectProduct("AA001")

	// err object should be null.
	assert.Nil(t, err)

	// product object should be AA001 product.
	assert.Equal(
		t,
		Product{
			ID:     "AA001",
			NameTH: "เตาย่าง หมายเลข 1",
			NameEN: "Thai stove No.1",
		},
		product,
		"Output should be AA001 product",
	)
}

// Product not found case of selectProduct function.
func TestFilterProductNotFound(t *testing.T) {
	products := Products{
		Product{
			ID:     "AA001",
			NameTH: "เตาย่าง หมายเลข 1",
			NameEN: "Thai stove No.1",
		},
		Product{
			ID:     "AA002",
			NameTH: "เตาย่าง หมายเลข 2",
			NameEN: "Thai stove No.2",
		},
		Product{
			ID:     "AA003",
			NameTH: "เตาย่าง หมายเลข 3",
			NameEN: "Thai stove No.3",
		},
	}

	product, err := products.selectProduct("AA005")

	// err object should not be null.
	assert.NotNil(t, err)

	// product object should be empty product struct.
	assert.Equal(
		t,
		Product{},
		product,
		"Output should be a emtry struct",
	)
}
