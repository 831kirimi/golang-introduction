package myimage

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func checkImageFormat(format string) bool {
	if format == "jpg" || format == "jpeg" || format == "png" || format == "gif" {
		return true
	} else {
		return false
	}
}

func getFormat(path string) (string, string) {
	iDot := strings.LastIndex(path, ".")
	return path[:iDot], path[iDot+1:]
}
func Conversion() {
	to := flag.String("t", "png", "image format option")
	flag.Parse()
	for _, dir := range flag.Args() {
		for _, file := range dirwalk(dir) {
			if name, format := getFormat(file); checkImageFormat(format) {
				var err error
				in, err := os.Open(file)
				if err != nil {
					panic(err)
				}
				defer in.Close()

				var img image.Image
				switch format {
				case "jpeg", "jpg":
					img, err = jpeg.Decode(in)
				case "png":
					img, err = png.Decode(in)
				case "gif":
					img, err = gif.Decode(in)
				default:
					fmt.Printf("format %s of %s is not supported\n", format, file)
				}
				if err != nil {
					panic(err)
				}

				out, err := os.Create(name + "." + *to)
				if err != nil {
					panic(err)
				}
				defer out.Close()

				switch *to {
				case "jpeg", "jpg":
					options := &jpeg.Options{
						Quality: 100,
					}
					err = jpeg.Encode(out, img, options)
				case "png":
					err = png.Encode(out, img)
				case "gif":
					options := &gif.Options{
						NumColors: 256,
						Quantizer: nil,
						Drawer:    nil,
					}
					err = gif.Encode(out, img, options)
				default:
					err = fmt.Errorf("format %s of %s is not supported", format, file)
				}
				if err != nil {
					panic(err)
				} else {
					fmt.Printf("%s converted to %s.%s\n", file, name, *to)
				}
			}
		}
	}
}
