package main

import (
	"fmt"
	"io/ioutil"

	"github.com/openalpr/openalpr/src/bindings/go/openalpr"
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

	resultFromFilePath, err := alpr.RecognizeByFilePath("./car.png")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", resultFromFilePath)
	fmt.Printf("\n\n\n")

	imageBytes, err := ioutil.ReadFile("./car2.png")
	if err != nil {
		fmt.Println(err)
	}
	resultFromBlob, err := alpr.RecognizeByBlob(imageBytes)
	fmt.Printf("%+v\n", resultFromBlob)
}
