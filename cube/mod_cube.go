package cube

type Cube struct {
	Title  string
	Length float64
	Wide   float64
	Height float64
}

/* Surface area */
func (c Cube) SurfaceArea() float64 {
	return 2*c.Length*c.Wide + 2*c.Wide*c.Height + 2*c.Length*c.Height
}

/* Volume */
func (c Cube) Volume() float64 {
	return c.Length * c.Wide * c.Height
}
