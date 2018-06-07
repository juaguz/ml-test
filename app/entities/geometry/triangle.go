package geometry


type Triangle struct {
	v1,v2,v3 PointImpl
}

func (t Triangle) Perimeter() float64{
	v1v2 := t.v1.Distance(t.v2)
	v2v3 := t.v2.Distance(t.v3)
	v3v1 := t.v3.Distance(t.v1)
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
	denominator := (t.v2.Y - t.v3.Y)*(t.v1.X - t.v3.X) + (t.v3.X - t.v2.X)*(t.v1.Y - t.v3.Y)
	a := ((t.v2.Y - t.v3.Y)*(point.X - t.v3.X) + (t.v3.X - t.v2.X)*(point.Y - t.v3.Y)) / denominator
	b := ((t.v3.Y - t.v1.Y)*(point.X - t.v3.X) + (t.v1.X - t.v3.X)*(point.Y - t.v3.Y)) / denominator
	c := 1 - a -b
	return 0 <= a && a <= 1 && 0 <= b && b <= 1 && 0 <= c && c <= 1
}
