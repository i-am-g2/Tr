package tr

import "math"

type HitRecord struct {
	P         Vector
	Normal    Vector
	T         float64
	frontFace bool
	Mat       Material
}

func (h *HitRecord) setFaceNormal(r *Ray, outwardNormal *Vector) {
	h.frontFace = r.Dir.Dot(outwardNormal) < 0
	if h.frontFace {
		h.Normal = *outwardNormal
	} else {
		h.Normal = *outwardNormal.ConstMult(-1)
	}
}

// Hittable a
type Hittable interface {
	Hit(*Ray, float64, float64) (bool, *HitRecord)
}

// Sphere a
type Sphere struct {
	Center Vector
	Radius float64
	Mat    Material
}

// Hit a
func (s Sphere) Hit(r *Ray, tmin float64, tmax float64) (bool, *HitRecord) {
	oc := r.Orig.SubVec(&s.Center)
	a := r.Dir.Dot(&r.Dir)
	b := oc.Dot(&r.Dir)
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - a*c

	if discriminant > 0 {
		root := math.Sqrt(discriminant)

		temp := (-b - root) / a
		if temp < tmax && temp > tmin {
			t := temp
			p := r.At(temp)
			outwardNormal := p.SubVec(&s.Center).ConstMult(1.0 / s.Radius)

			tempRecord := &HitRecord{*p, *outwardNormal, t, false, s.Mat}
			tempRecord.setFaceNormal(r, outwardNormal)

			return true, tempRecord
		}

		temp = (-b + root) / a
		if temp < tmax && temp > tmin {
			t := temp
			p := r.At(temp)
			outwardNormal := p.SubVec(&s.Center).ConstMult(1.0 / s.Radius)
			tempRecord := &HitRecord{*p, *outwardNormal, t, false, s.Mat}
			tempRecord.setFaceNormal(r, outwardNormal)
			return true, tempRecord
		}
	}

	return false, nil

}

type HittableList struct {
	objects []Hittable
}

func (h *HittableList) Add(obj Hittable) {
	if h.objects == nil {
		h.objects = make([]Hittable, 0)
	}
	h.objects = append(h.objects, obj)
}

func (h *HittableList) Hit(r *Ray, tmin, tmax float64) (bool, *HitRecord) {
	hitAny := false
	var hitRecordClosest HitRecord
	closest := tmax

	for _, obj := range h.objects {
		tmp, tmpa := obj.Hit(r, tmin, closest)
		if tmp {
			hitAny = true
			hitRecordClosest = *tmpa
			closest = tmpa.T
		}
	}
	return hitAny, &hitRecordClosest
}

func NewSphere(center *Vector, radius float64, m Material) *Sphere {
	var sp Sphere
	sp.Center = *center
	sp.Radius = radius
	sp.Mat = m
	return &sp
}
