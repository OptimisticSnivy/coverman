package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

var files []string

func walk(s string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !d.IsDir() && filepath.Ext(s) == ".jpg" {
		file, _ := filepath.Abs(s)
		files = append(files, file)
	}
	return nil
}

func Map(x int, in_min int, in_max int, out_min int, out_max int) int {
	return ((x-in_min)*(out_max-out_min)/(in_max-in_min) + out_min)
}

func main() {
	arg := os.Args[1:]
	var pixels []int
	var density string = ".:-=+*#%@"

	filepath.WalkDir(arg[0], walk)

	f, err := os.Open(files[3])
	if err != nil {
		log.Fatal(err)
	}

	img, _, error := image.Decode(f)
	if error != nil {
		log.Fatal(error)
	}

	rf := resize.Resize(60, 60, img, resize.NearestNeighbor)
	for i := 0; i < rf.Bounds().Dx(); i++ {
		for j := 0; j < rf.Bounds().Dy(); j++ {
			r, g, b, _ := rf.At(j, i).RGBA()
			br := int(r>>8+g>>8+b>>8) / 3
			brm := Map(br, 0, 255, 0, 8)
			pixels = append(pixels, brm)
			fmt.Print(string(density[brm]), string(density[brm]))
		}
		fmt.Print("\n")
	}

	resized, _ := os.Create("resized.jpg")

	jpeg.Encode(resized, rf, nil)
	defer f.Close()
}
