package geometry


type Triangle struct {
	V1,V2, V3 PointImpl
}

//Calcula el perimetro del triangulo
//Obtiene la distancia entre los vertices del triangulo
//y realiza la suma
func (t Triangle) Perimeter() float64{
	v1v2 := t.V1.Distance(t.V2)
	v2v3 := t.V2.Distance(t.V3)
	v3v1 := t.V3.Distance(t.V1)
	return v1v2+v2v3+v3v1
}


//Calcula si un punto se encuentra dentro de un triangulo
func (t Triangle) PointInside(point PointImpl) bool {
	denominator := (t.V2.Y - t.V3.Y)*(t.V1.X - t.V3.X) + (t.V3.X - t.V2.X)*(t.V1.Y - t.V3.Y)
	a := ((t.V2.Y - t.V3.Y)*(point.X - t.V3.X) + (t.V3.X - t.V2.X)*(point.Y - t.V3.Y)) / denominator
	b := ((t.V3.Y - t.V1.Y)*(point.X - t.V3.X) + (t.V1.X - t.V3.X)*(point.Y - t.V3.Y)) / denominator
	c := 1 - a -b
	return 0 <= a && a <= 1 && 0 <= b && b <= 1 && 0 <= c && c <= 1
}
