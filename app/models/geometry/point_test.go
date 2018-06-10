package geometry

import (
	"testing"
	"fmt"
)

func TestPointImplDistance(t *testing.T) {
	p1 := PointImpl{-3, 1}
	p2 := PointImpl{6, 4}

	distance := p1.Distance(p2)

	if distance < 9.48 || distance > 9.49 {
		t.Fatal(fmt.Sprintf("result should be between 9.48 and 9.49 and its %v", distance))
	}
}


func TestInline(t *testing.T){
	p1 := PointImpl{-3, 1}
	p2 := PointImpl{0, 2}
	p3 := PointImpl{6, 4}
	isInline := p1.IsCollinear(p2,p3)
	if !isInline {
		t.Fatalf("result should  true and  not %v", isInline)
	}

}

func TestPointImplIsInline(t *testing.T) {
	p1 := PointImpl{-2, -1}
	p2 := PointImpl{0, 0}
	p3 := PointImpl{4, 2}
	p4 := PointImpl{6, 3}
	p5 := PointImpl{8, 4}
	isInline := p1.IsCollinear(p2,p3, p4,p5)
	if !isInline {
		t.Fatalf("result should  true and  not %v", isInline)
	}
}