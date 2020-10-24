package tr

type Ray struct {
	Orig Vector
	Dir  Vector
}

func (r *Ray) At(t float64) *Vector {
	temp := r.Dir.ConstMult(t)
	return r.Orig.AddVec(temp)
}

func NewRay(orig, dir *Vector) *Ray {
	return &Ray{*orig, *dir}
}
