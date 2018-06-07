package entities


type Triangle struct {
	v1,v2,v3 PointImpl
}

func (t Triangle) Perimeter() float64{
	v1v2 := t.v1.Distance(t.v2)
	v2v3 := t.v2.Distance(t.v3)
	v3v1 := t.v3.Distance(t.v1)
	return v1v2+v2v3+v3v1
}