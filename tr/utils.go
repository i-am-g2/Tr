package tr

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

func UnitVec(r *Vector) *Vector {
	d := math.Sqrt(r.X*r.X + r.Y*r.Y + r.Z*r.Z)
	temp := Vector{r.X / d, r.Y / d, r.Z / d}
	return &temp

}

func Cross(u, v *Vector) *Vector {
	return NewVector(
		u.Y*v.Z-u.Z*v.Y,
		u.Z*v.X-u.X*v.Z,
		u.X*v.Y-u.Y*v.X,
	)
}

func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func DegToRad(x float64) float64 {
	return x / 180 * math.Pi
}

func RandomMinMax(min, max float64) float64 {
	return (rand.Float64() * (max - min)) + min
}

func ProgressBar(done, total int) {
	scale := 50.0
	fmt.Fprintf(os.Stderr, "Rendering [")
	percentDone := int((float64(done) / float64(total)) * scale)
	for i := 0; i < percentDone; i++ {
		fmt.Fprintf(os.Stderr, "=")
	}
	for i := percentDone; i < int(scale); i++ {
		fmt.Fprintf(os.Stderr, " ")
	}
	fmt.Fprintf(os.Stderr, "] %d out of %d \r", done+1, total)
	if done+1 == total {
		fmt.Fprintf(os.Stderr, "\n")
	}
}

func Log(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
}
