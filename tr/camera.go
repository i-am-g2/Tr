package tr

import "math"

type Camera struct {
	AspectRatio float64
	Vfov        float64

	theta      float64
	h          float64
	focusDist  float64
	apperture  float64
	lensRadius float64

	W, U, V *Vector

	//Camera
	ViewportHeight float64
	ViewportWidth  float64

	Origin          *Vector
	Horizontal      *Vector
	Vertical        *Vector
	LowerLeftCorner *Vector
}

func (c *Camera) GetRay(u, v float64) *Ray {
	rd := RandomUnitVector().ConstMult(c.lensRadius)
	offset := u*rd.X + v*rd.Y
	dir := c.LowerLeftCorner.AddVec(c.Horizontal.ConstMult(u)).AddVec(c.Vertical.ConstMult(v)).SubVec(c.Origin).SubVec(NewVector(offset, offset, offset))
	r := NewRay(c.Origin.AddVec(NewVector(offset, offset, offset)), dir)
	return r
}

func InitCamera(lookFrom, lookAt, vup *Vector, AspectRatio, Vfov, apperture, focusDist float64) *Camera {

	var cam Camera
	cam.AspectRatio = AspectRatio
	cam.theta = DegToRad(Vfov)

	cam.h = math.Tan(cam.theta / 2)

	//Camera
	cam.ViewportHeight = 2.0 * cam.h
	cam.ViewportWidth = cam.AspectRatio * cam.ViewportHeight

	cam.W = UnitVec(lookFrom.SubVec(lookAt))
	cam.U = Cross(vup, cam.W)
	cam.V = Cross(cam.W, cam.U)

	cam.lensRadius = apperture / 2

	cam.Origin = lookFrom
	cam.Horizontal = cam.U.ConstMult(cam.ViewportWidth).ConstMult(focusDist)
	cam.Vertical = cam.V.ConstMult(cam.ViewportHeight).ConstMult(focusDist)
	cam.LowerLeftCorner = cam.Origin.SubVec(cam.Horizontal.ConstMult(0.5)).SubVec(cam.Vertical.ConstMult(0.5)).SubVec(cam.W.ConstMult(focusDist))

	return &cam
}
