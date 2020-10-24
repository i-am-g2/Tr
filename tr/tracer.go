package tr

import "math"

func RayColor(r *Ray, w *HittableList, depth int) *Vector {
	hit, rec := w.Hit(r, 0.001, math.Inf(0))
	if depth <= 0 {
		return NewVector(0, 0, 0)
	}
	if hit {
		b, scatter, atten := rec.Mat.Scatter(r, rec)
		if b {
			return atten.Hamadard(RayColor(scatter, w, depth-1))
		}
		return NewVector(0, 0, 0)
	}

	unitDir := UnitVec(&r.Dir)
	t := 0.5 * (unitDir.Y + 1.0)
	temp := NewVector(1.0, 1.0, 1.0).ConstMult(1.0 - t).AddVec(NewVector(0.5, 0.7, 1.0).ConstMult(t))
	return temp
}
