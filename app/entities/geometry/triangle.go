package geometry


type Triangle struct {
	V1,V2, V3 PointImpl
}

func (t Triangle) Perimeter() float64{
	v1v2 := t.V1.Distance(t.V2)
	v2v3 := t.V2.Distance(t.V3)
	v3v1 := t.V3.Distance(t.V1)
	return v1v2+v2v3+v3v1
}


//function pointInTriangle(x1, y1, x2, y2, x3, y3, x, y:Number):Boolean
//{
//var denominator:Number = ((y2 - y3)*(x1 - x3) + (x3 - x2)*(y1 - y3));
//var a:Number = ((y2 - y3)*(x - x3) + (x3 - x2)*(y - y3)) / denominator;
//var b:Number = ((y3 - y1)*(x - x3) + (x1 - x3)*(y - y3)) / denominator;
//var c:Number = 1 - a - b;
//
//return 0 <= a && a <= 1 && 0 <= b && b <= 1 && 0 <= c && c <= 1;
//}

func (t Triangle) PointInside(point PointImpl) bool {
	denominator := (t.V2.Y - t.V3.Y)*(t.V1.X - t.V3.X) + (t.V3.X - t.V2.X)*(t.V1.Y - t.V3.Y)
	a := ((t.V2.Y - t.V3.Y)*(point.X - t.V3.X) + (t.V3.X - t.V2.X)*(point.Y - t.V3.Y)) / denominator
	b := ((t.V3.Y - t.V1.Y)*(point.X - t.V3.X) + (t.V1.X - t.V3.X)*(point.Y - t.V3.Y)) / denominator
	c := 1 - a -b
	return 0 <= a && a <= 1 && 0 <= b && b <= 1 && 0 <= c && c <= 1
}
