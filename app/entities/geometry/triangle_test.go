package geometry

import "testing"

func TestTrianglePerimeter(t *testing.T) {
	v1 := PointImpl{1,1}
	v2 := PointImpl{3,5}
	v3 := PointImpl{4,2}
	triangle := Triangle{v1,v2,v3}
	result := triangle.Perimeter()
	if result < 10.79 || result > 10.80 {
		t.Fatalf("result should be between 10.79 and 10.80 not %v", result)
	}
}


func TestTrianglePointInside(t *testing.T) {
	v1 := PointImpl{1,1}
	v2 := PointImpl{3,5}
	v3 := PointImpl{4,2}
	triangle := Triangle{v1,v2,v3}
	result := triangle.PointInside(PointImpl{2,2})
	if !result {
		t.Fatalf("result should be true not %v", result)
	}
}


func TestTrianglePointNotInside(t *testing.T) {
	v1 := PointImpl{1,1}
	v2 := PointImpl{3,5}
	v3 := PointImpl{4,2}
	triangle := Triangle{v1,v2,v3}
	result := triangle.PointInside(PointImpl{0,0})
	if result {
		t.Fatalf("result should be false not %v", result)
	}
}


