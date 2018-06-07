package utils

import (
	"testing"

)

func TestDegreeToRadian(t *testing.T) {
	var degree float64
	degree = 30
	result := DegreeToRadian(degree)

	if result < float64(0.52)  || result > float64(0.53) {
		t.Fatalf("result should be between 0.52 and 0.53 and its %v", result)
	}
}




