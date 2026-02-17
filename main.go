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

func main() {
	arg := os.Args[1:]

	filepath.WalkDir(arg[0], walk)
	fmt.Println(files[0])

	f, err := os.Open(files[2])
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(f)
	r := resize.Resize(60, 0, img, resize.Lanczos3)

	resized, err := os.Create("resized.jpg")
	if err != nil {
		log.Fatal(err)
	}

	jpeg.Encode(resized, r, nil)
	defer f.Close()
}
