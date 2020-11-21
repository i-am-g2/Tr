package main

import (
	"image"
	"image/png"
	"math/rand"
	"os"
	"sync"

	"github.com/i-am-g2/tr/tr"
)

func main() {

	// Setting up Camera Parameters
	aspectRatio := 3.0 / 2.0
	lookFrom := tr.NewVector(13, 2, 3)
	lookAt := tr.NewVector(0, 0, 0)
	vup := tr.NewVector(0, 1, 0)
	dist := 10.0
	apperture := 0.1
	cam := tr.InitCamera(lookFrom, lookAt, vup, aspectRatio, 20.0, apperture, dist)

	// Setting Up Image Parameters
	imageWidth := 1200
	imageHeight := int(float64(imageWidth) / aspectRatio)
	samplePerPixel := 100
	maxDepth := 50
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{imageWidth, imageHeight}})

	// Creating Random Scene
	var world tr.HittableList
	world = *randomScene()

	// Ray Tracing Start
	var wg sync.WaitGroup
	for j := 0; j < imageHeight; j++ {
		tr.ProgressBar(j, imageHeight)
		for i := 0; i < imageWidth; i++ {
			i, j := i, j
			// This is Important, By default same instance of variable is used in loop. So while one go func capture
			// the i,j and rendering,  the value of i,j is changed as loop proceeds
			wg.Add(1)
			go func() {
				color := tr.NewVector(0, 0, 0)
				for s := 0; s < samplePerPixel; s++ {
					u := (float64(i) + rand.Float64()) / float64(imageWidth-1)
					v := (float64(j) + rand.Float64()) / float64(imageHeight-1)

					r := cam.GetRay(u, v)
					color = color.AddVec(tr.RayColor(r, &world, maxDepth))
				}

				color.SetColor(i, j, samplePerPixel, img)
				wg.Done()
			}()
		}
		wg.Wait()
	}

	pngFile, _ := os.Create("outputs/result.png")
	png.Encode(pngFile, img)
	pngFile.Close()

}

func randomScene() *tr.HittableList {
	var world tr.HittableList
	groundMaterial := tr.NewLambertian(tr.NewVector(0.5, 0.5, 0.5))

	world.Add(tr.NewSphere(tr.NewVector(0, -1000, 0), 1000, groundMaterial))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := tr.NewVector(float64(a)+0.9*rand.Float64(), 0.2, float64(b)+0.9*rand.Float64())

			if center.SubVec(tr.NewVector(4, 0.2, 0)).LengthSquared() > 0.9*0.9 {

				if chooseMat < 0.8 {

					color := tr.NewRandom(0, 1).Hamadard(tr.NewRandom(0, 1))
					mat := tr.NewLambertian(color)
					world.Add(tr.NewSphere(center, 0.2, mat))

				} else if chooseMat < 0.95 {

					color := tr.NewRandom(0, 0.5) // Random 3 dimensional vector with values between 0 to 0.5
					fuzziness := tr.RandomMinMax(0, 0.5)
					mat := tr.NewMetal(color, fuzziness)
					world.Add(tr.NewSphere(center, 0.2, mat))

				} else {
					mat := tr.NewDielectric(1.5)
					world.Add(tr.NewSphere(center, 0.2, mat))
				}

			}

		}
	}
	mat := tr.NewDielectric(1.5)
	world.Add(tr.NewSphere(tr.NewVector(0, 1, 0), 1.0, mat))

	mat2 := tr.NewLambertian(tr.NewVector(0.4, 0.2, 0.1))
	world.Add(tr.NewSphere(tr.NewVector(-4, 1, 0), 1.0, mat2))

	mat3 := tr.NewMetal(tr.NewVector(0.7, 0.6, 0.5), 0.0)
	world.Add(tr.NewSphere(tr.NewVector(4, 1, 0), 1.0, mat3))

	return &world
}
