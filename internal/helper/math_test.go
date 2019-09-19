package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomInteger(t *testing.T) {
	sliceLength := 5
	output := GenerateRandomInteger(sliceLength)

	// Output length should equal to sliceLength input.
	assert.Equal(
		t,
		len(output),
		sliceLength,
		fmt.Sprintf("Slice length should be %d", sliceLength),
	)
}
