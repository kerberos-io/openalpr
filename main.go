package main

import (
	"fmt"
	"io/ioutil"

	"github.com/openalpr/openalpr/src/bindings/go/openalpr"
	"gocv.io/x/gocv"
)

func main() {
	alpr := openalpr.NewAlpr("eu", "", "/var/lib/openalpr/runtime_data")
	defer alpr.Unload()

	if !alpr.IsLoaded() {
		fmt.Println("OpenAlpr failed to load!")
		return
	}
	alpr.SetTopN(20)

	fmt.Println(alpr.IsLoaded())
	fmt.Println(openalpr.GetVersion())

	fileWithLicensePlate := "./car2.png"
	imageBytes, err := ioutil.ReadFile(fileWithLicensePlate)
	if err != nil {
		fmt.Println(err)
	}
	resultFromBlob, err := alpr.RecognizeByBlob(imageBytes)
	fmt.Printf("%+v\n", resultFromBlob)

	// Read image to Mat
	mat := gocv.IMRead(fileWithLicensePlate, gocv.IMReadColor)
	fmt.Println(mat)
}
