package tr

import (
	"fmt"
	"math"
	"os"
)

type Vector struct {
	X, Y, Z float64
}

func (v *Vector) WriteColor(samplePerPixel int) {

	scale := 1.0 / float64(samplePerPixel)

	fmt.Fprintln(os.Stdout, int(256*Clamp(math.Sqrt(v.X*scale), 0.0, 0.9999)), int(256*Clamp(math.Sqrt(v.Y*scale), 0.0, 0.9999)), int(256*Clamp(math.Sqrt(v.Z*scale), 0.0, 0.9999)))

}

func (v *Vector) ConstMult(t float64) *Vector {
	return &Vector{v.X * t, v.Y * t, v.Z * t}
}

// type Color Vec3
func (v *Vector) AddVec(t *Vector) *Vector {
	return &Vector{v.X + t.X, v.Y + t.Y, v.Z + t.Z}
}
func (v *Vector) SubVec(t *Vector) *Vector {
	return &Vector{v.X - t.X, v.Y - t.Y, v.Z - t.Z}
}

func NewVector(x, y, z float64) *Vector {
	return &Vector{x, y, z}
}

func (v *Vector) ToString() string {
	return fmt.Sprintf("%.6f %0.6f + %0.6f", v.X, v.Y, v.Z)
}

func (v *Vector) Dot(t *Vector) float64 {
	return v.X*t.X + v.Y*t.Y + v.Z*t.Z
}

func (v *Vector) LengthSquared() float64 {
	return v.Dot(v)
}

func RandomUnitVector() *Vector {
	a := RandomMinMax(0, 2*math.Pi)
	z := RandomMinMax(-1, 1)
	r := math.Sqrt(1 - z*z)
	return NewVector(r*math.Cos(a), r*math.Sin(a), z)
}

func (v *Vector) Hamadard(t *Vector) *Vector {
	return NewVector(v.X*t.X, v.Y*t.Y, v.Z*t.Z)
}

func NewRandom(x, y float64) *Vector {
	return &Vector{RandomMinMax(x, y), RandomMinMax(x, y), RandomMinMax(x, y)}
}
