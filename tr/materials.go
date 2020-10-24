package tr

import (
	"math"
	"math/rand"
)

type Material interface {
	Scatter(*Ray, *HitRecord) (bool, *Ray, *Vector)
}

// Lambertian Start
type Lambertian struct {
	Albedo Vector
}

func NewLambertian(v *Vector) *Lambertian {
	var material Lambertian
	material.Albedo = *v
	return &material
}

func (l Lambertian) Scatter(r *Ray, rec *HitRecord) (bool, *Ray, *Vector) {
	scatterDir := rec.Normal.AddVec(RandomUnitVector())
	scattered := NewRay(&rec.P, scatterDir)
	attenuation := l.Albedo
	return true, scattered, &attenuation
}

type Metal struct {
	Albedo Vector
	Fuzz   float64
}

func NewMetal(v *Vector, fuzz float64) *Metal {
	var material Metal
	material.Albedo = *v
	material.Fuzz = fuzz
	return &material

}
func (m Metal) Scatter(r *Ray, rec *HitRecord) (bool, *Ray, *Vector) {
	reflected := Reflect(UnitVec(&r.Dir), &rec.Normal)
	scattered := NewRay(&rec.P, reflected.AddVec(RandomUnitVector().ConstMult(m.Fuzz)))
	return scattered.Dir.Dot(&rec.Normal) > 0, scattered, &m.Albedo
}

type Dielectric struct {
	Ir float64
}

func NewDielectric(ir float64) *Dielectric {
	var material Dielectric
	material.Ir = ir
	return &material
}

func (m Dielectric) Scatter(r *Ray, rec *HitRecord) (bool, *Ray, *Vector) {
	attenuation := NewVector(1.0, 1.0, 1.0)
	refractionRatio := m.Ir
	if rec.frontFace == true {
		refractionRatio = 1.0 / m.Ir
	}

	unitDir := UnitVec(&r.Dir)
	cosTheta := math.Min(rec.Normal.Dot(unitDir.ConstMult(-1)), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	cannotRefract := refractionRatio*sinTheta > 1.0

	dir := NewVector(0, 0, 0)
	if cannotRefract || (Reflectance(cosTheta, refractionRatio) > rand.Float64()) {
		dir = Reflect(unitDir, &rec.Normal)
	} else {
		dir = Refract(unitDir, &rec.Normal, refractionRatio)
	}
	scattered := NewRay(&rec.P, dir)
	return true, scattered, attenuation

}

func Reflectance(cosine, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}

func Reflect(V, N *Vector) *Vector {
	return V.SubVec(N.ConstMult(2 * V.Dot(N)))
}

func Refract(uv, n *Vector, et float64) *Vector {
	cosTheta := uv.ConstMult(-1).Dot(n)
	routperp := uv.AddVec(n.ConstMult(cosTheta)).ConstMult(et)
	routparallel := n.ConstMult(-1 * math.Sqrt(math.Abs(1.0-routperp.LengthSquared())))
	return routparallel.AddVec(routperp)

}
