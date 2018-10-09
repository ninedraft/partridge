package main

import (
	"fmt"
	"image/color"
	"math"
	. "misc/partridge/vec"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	var users = []string{
		"phillennium",
		"olegchir",
		"ragequit",
		"mi5ha6in",
		"Milfgard",
		"HotWaterMusic",
		"Jeditobe",
		"it_man",
		"AloneCoder",
		"MagisterLudi",
	}
	var raitings = Vector{
		1058.6,
		// 5000,
		733.0,
		710.7,
		689.9,
		629.2,
		620.9,
		546.9,
		524.3,
		500.9,
		495.4,
	}
	var mean = raitings.Mean()
	var median = raitings.Median()
	fmt.Printf("mean   %v\nmedian %v\n", mean, median)
	// -> 650.9799999999999

	var line, errLine = plotter.NewLine(NewPoints(Vector{-0.4, 9.3}, Vector{mean, mean}))
	if errLine != nil {
		panic(errLine)
	}
	line.Color = color.RGBA{R: 255, A: 255}
	canvas, err := plot.New()
	if err != nil {
		panic(err)
	}

	canvas.Title.Text = "Рейтинг пользователей Хабра"
	canvas.Y.Label.Text = "Рейтинг"
	var bar, errNewScatter = plotter.NewBarChart(raitings, 12*vg.Millimeter)
	if errNewScatter != nil {
		panic(errNewScatter)
	}
	bar.Color = color.RGBA{G: 255, A: 255}

	var medianLevel, errMedianLevel = plotter.NewLine(NewPoints(Vector{-0.4, 9.3}, Vector{median, median}))
	if errMedianLevel != nil {
		panic(errMedianLevel)
	}
	medianLevel.Color = color.RGBA{B: 255, A: 255}
	canvas.NominalX(users...)
	canvas.X.Label.YAlign = draw.YCenter
	canvas.Add(bar, medianLevel)
	canvas.Legend.Add("Рейтинг", bar)
	canvas.Legend.Add("Средний рейтинг", line)
	canvas.Legend.Add("Медиана", medianLevel)
	canvas.X.Tick.Label.Rotation = math.Pi / 4
	//	canvas.Legend.Add("Bar", bar)
	canvas.Legend.Top = true
	canvas.Legend.Left = false

	canvas.Add(line)
	if err := canvas.Save(20*vg.Centimeter, 12*vg.Centimeter, "median_ratings.png"); err != nil {
		panic(err)
	}
}
