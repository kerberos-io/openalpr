package main

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"

	"github.com/openalpr/openalpr/src/bindings/go/openalpr"
	"gocv.io/x/gocv"
)

func main() {
	alpr := openalpr.NewAlpr("eu", "/etc/openalpr/openalpr.conf", "/var/lib/openalpr/runtime_data")

	defer alpr.Unload()

	if !alpr.IsLoaded() {
		fmt.Println("OpenAlpr failed to load!")
		return
	}
	alpr.SetTopN(20)

	fileWithLicensePlate := "./car8.png"
	imageBytes, err := ioutil.ReadFile(fileWithLicensePlate)
	if err != nil {
		fmt.Println(err)
	}
	resultFromBlob, err := alpr.RecognizeByBlob(imageBytes)
	fmt.Printf("%+v\n", resultFromBlob)

	// Plates
	plates := resultFromBlob.Plates

	// Get first plate, we only assume 1 plate.
	plate := plates[0]
	// Get plate number
	platePoints := plate.PlatePoints

	// Read image to Mat
	mat := gocv.IMRead(fileWithLicensePlate, gocv.IMReadColor)

	// Convert plate coordinates to rectangle
	if len(platePoints) != 4 {
		fmt.Println("Plate points are not 4")
		return
	}

	x := platePoints[0].X
	y := platePoints[0].Y
	w := platePoints[2].X - platePoints[0].X
	h := platePoints[2].Y - platePoints[0].Y

	// Rectangle
	rect := image.Rect(x, y, x+w, y+h)
	// Draw rectangle
	gocv.Rectangle(&mat, rect, color.RGBA{0, 255, 0, 0}, 2)

	// Get best match plate number
	plateNumber := plate.TopNPlates[0].Characters
	fmt.Println("License Plate: " + plateNumber)

	// Draw text
	gocv.PutText(&mat, plateNumber, image.Pt(x, y-20), gocv.FontHersheyPlain, 2, color.RGBA{0, 255, 0, 0}, 2)

	// Save image to tmp folder
	gocv.IMWrite("/tmp/car8_out.png", mat)
}
