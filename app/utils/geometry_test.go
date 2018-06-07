package utils

import (
	"testing"
	"fmt"
)

func TestDegreeToRadian(t *testing.T) {
	var degree float64
	degree = 30
	result := DegreeToRadian(degree)

	if result < float64(0.52)  || result > float64(0.53) {
		t.Fatal(fmt.Sprintf("result should be between 0.52 and 0.53 and its %v", result))
	}
}

