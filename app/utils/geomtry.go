package utils

import "math"

func DegreeToRadian(degree float64) float64{
	return (degree / float64(180)) * math.Pi
}


