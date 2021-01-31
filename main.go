package main

import (
	"fmt"
	"github.com/AllenDang/giu"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
)
// HLine draws a horizontal line
func HLine(x1, y, x2 int) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// VLine draws a veritcal line
func VLine(x, y1, y2 int) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(x1, y1, x2, y2 int) {
	HLine(x1, y1, x2)
	HLine(x1, y2, x2)
	VLine(x1, y1, y2)
	VLine(x2, y1, y2)
}

func RectFill(x1, y1, x2, y2 int) {
	for ; x1 < x2; x1++ {
		VLine(x1, y1, y2)
	}
}

func DrawCross(x1, y1 int, param ...int) {
	size := 3
	if len(param) > 0 {
		size = param[0]
	}
	VLine(x1, y1-size, y1+size)
	HLine(x1-size, y1, x1+size)
}

var width = 2480
var height = 3508
var img = image.NewRGBA(image.Rect(0, 0, width, height))
var col color.Color

type Point struct {
	x int
	y int
}

func RandPoint(width, height int) Point {
	return Point{
		x : rand.Intn(width),
		y:rand.Intn(height),
	}
}

func Dist(p1 Point, p2 Point) float64  {
	x := (p1.x - p2.x) * (p1.x - p2.x)
	y := (p1.y - p2.y) * (p1.y - p2.y)
	return math.Sqrt(float64(x) + float64(y))
}

func IsNearPoints(p1 Point, pts []*Point, dis float64) bool {
	for _, p2 := range pts {
		if Dist(p1, *p2) < dis {
			return true
		}
	}
	return false
}
func genRandomDots(){
	col = color.RGBA{255, 255, 255, 255}
	RectFill(0, 0, width,  height)
	col = color.RGBA{0, 0, 0, 255}
	pts := make([]*Point, 0)
	j := 1000
	for i := 0; i < 100; {
		newP := RandPoint(width, height)
		if IsNearPoints(newP, pts, 200) {
			j -= 1
			if j < 0 {
				fmt.Printf("too much invalid points found, end with %d points\n", i)
				return
			}
			continue
		}
		i++
		pts = append(pts, &newP)
		DrawCross(newP.x, newP.y)
	}

	f, err := os.Create("draw.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

func loop() {
	giu.SingleWindow("DHT").Layout(
		giu.Label("Generate:"),
		giu.Button("Random dots").OnClick(genRandomDots),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Drawer Helper tool", 400, 200, giu.MasterWindowFlagsNotResizable, nil)
	wnd.Run(loop)
}