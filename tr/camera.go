package tr

import "math"

type Camera struct {
	lensRadius float64

	W, U, V         *Vector
	Origin          *Vector
	Horizontal      *Vector
	Vertical        *Vector
	LowerLeftCorner *Vector
}

func (c *Camera) GetRay(u, v float64) *Ray {
	rd := RandomUnitVector().ConstMult(c.lensRadius)
	offset := u*rd.X + v*rd.Y

	// dir = lowerLeft + u* Horizontal + v * Vertical - origin - offset
	dir := c.LowerLeftCorner.AddVec(c.Horizontal.ConstMult(u)).AddVec(c.Vertical.ConstMult(v)).SubVec(c.Origin).SubTrip(offset, offset, offset)
	r := NewRay(c.Origin.AddTrip(offset, offset, offset), dir)
	return r
}

func InitCamera(lookFrom, lookAt, vup *Vector, AspectRatio, Vfov, apperture, focusDist float64) *Camera {

	var cam Camera
	theta := DegToRad(Vfov)

	h := math.Tan(theta / 2)

	//Camera
	ViewportHeight := 2.0 * h
	ViewportWidth := AspectRatio * ViewportHeight

	cam.W = UnitVec(lookFrom.SubVec(lookAt))
	cam.U = Cross(vup, cam.W)
	cam.V = Cross(cam.W, cam.U)

	cam.lensRadius = apperture / 2

	cam.Origin = lookFrom
	cam.Horizontal = cam.U.ConstMult(ViewportWidth).ConstMult(focusDist)
	cam.Vertical = cam.V.ConstMult(ViewportHeight).ConstMult(focusDist)
	cam.LowerLeftCorner = cam.Origin.SubVec(cam.Horizontal.ConstMult(0.5)).SubVec(cam.Vertical.ConstMult(0.5)).SubVec(cam.W.ConstMult(focusDist))

	return &cam
}
