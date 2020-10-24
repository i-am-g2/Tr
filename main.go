package main

import (
	"fmt"
	"math/rand"
	"os"

	. "github.com/i-am-g2/tr/tr"
)

func main() {
	// Image

	AspectRatio := 3.0 / 2.0
	lookFrom := NewVector(13, 2, 3)
	lookAt := NewVector(0, 0, 0)
	vup := NewVector(0, 1, 0)
	dist := 10.0
	apperture := 0.1
	cam := InitCamera(lookFrom, lookAt, vup, AspectRatio, 20.0, apperture, dist)

	imageWidth := 200
	imageHeight := int(float64(imageWidth) / cam.AspectRatio)
	samplePerPixel := 100
	maxDepth := 50
	var world HittableList
	world = *randomScene()

	fmt.Fprintln(os.Stdout, "P3")
	fmt.Fprintln(os.Stdout, imageWidth, imageHeight)
	fmt.Fprintln(os.Stdout, "255")

	for j := imageHeight - 1; j >= 0; j-- {
		ProgressBar(imageHeight-j-1, imageHeight)
		for i := 0; i < imageWidth; i++ {
			color := NewVector(0, 0, 0)
			for s := 0; s < samplePerPixel; s++ {

				u := (float64(i) + rand.Float64()) / float64(imageWidth-1)
				v := (float64(j) + rand.Float64()) / float64(imageHeight-1)

				r := cam.GetRay(u, v)

				color = color.AddVec(RayColor(r, &world, maxDepth))
			}
			color.WriteColor(samplePerPixel)

		}
	}
}

func randomScene() *HittableList {
	var world HittableList
	groundMaterial := Lambertian{*NewVector(0.5, 0.5, 0.5)}

	world.Add(NewSphere(NewVector(0, -1000, 0), 1000, groundMaterial))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := NewVector(float64(a)+0.9*rand.Float64(), 0.2, float64(b)+0.9*rand.Float64())

			if center.SubVec(NewVector(4, 0.2, 0)).LengthSquared() > 0.9*0.9 {

				if chooseMat < 0.8 {
					color := NewRandom(0, 1).Hamadard(NewRandom(0, 1))
					mat := Lambertian{*color}
					world.Add(NewSphere(center, 0.2, mat))
				} else if chooseMat < 0.95 {
					mat := Metal{*NewRandom(0, 0.5), RandomMinMax(0, 0.5)}
					world.Add(NewSphere(center, 0.2, mat))

				} else {
					mat := Dielectric{1.5}
					world.Add(NewSphere(center, 0.2, mat))
				}

			}

		}
	}
	mat := Dielectric{1.5}
	world.Add(NewSphere(NewVector(0, 1, 0), 1.0, mat))

	mat2 := Lambertian{*NewVector(0.4, 0.2, 0.1)}
	world.Add(NewSphere(NewVector(-4, 1, 0), 1.0, mat2))

	mat3 := Metal{*NewVector(0.7, 0.6, 0.5), 0.0}
	world.Add(NewSphere(NewVector(4, 1, 0), 1.0, mat3))

	return &world
}
