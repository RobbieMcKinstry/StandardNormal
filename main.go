package main

import (
	. "github.com/RobbieMcKinstry/StandardNormal/stdnormal"
	"github.com/RobbieMckinstry/UniformRandom/RandGenerator"
	"github.com/RobbieMckinstry/UniformRandom/StatTests"
	plotLib "github.com/gonum/plot"
	. "github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"

	"image/color"
	"log"
	"math"
	"math/rand"
	"time"
)

const (
	Executions = 10000000 // ten million
)

var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StdNorm(x float64) float64 {
	y := (2.0 / math.Sqrt(2.0*math.Pi)) * math.Exp(-2.0*x*x)
	return y
}

func main() {

	inverseData := make(XYs, Executions)
	acceptData := make(XYs, Executions)
	polarData := make(XYs, Executions)

	// Generate the inverse transform numbers
	for i := 0; i < Executions; i++ {
		x := InverseTransform(r)
		y := StdNorm(x)

		point := struct {
			X float64
			Y float64
		}{x, y}
		inverseData = append(inverseData, point)
	}

	// Generate the histogram
	plot, err := plotLib.New()
	if err != nil {
		log.Fatal(err)
	}
	plot.Title.Text = "Inverse Transform Outputs"
	h, err := NewHistogram(inverseData, 50)
	h.Width = 0.1
	h.Color = color.RGBA{B: 255, A: 255}

	if err != nil {
		log.Panic(err)
	}
	h.Normalize(1)
	plot.Add(h)

	// The normal distribution function
	norm := NewFunction(StdNorm)
	norm.Color = color.RGBA{R: 255, A: 255}
	norm.Width = vg.Points(2)
	plot.Add(norm)

	err = plot.Save(400, 400, "inverse_transform_histogram.png")
	if err != nil {
		log.Panic(err)
	}

	// Generate the accept/reject numbers
	for i := 0; i < Executions; i++ {
		x, _ := AcceptReject(r)
		y := StdNorm(x)
		point := struct {
			X float64
			Y float64
		}{x, y}

		acceptData = append(acceptData, point)
	}

	// Generate the histogram
	plot, err = plotLib.New()
	if err != nil {
		log.Fatal(err)
	}
	plot.Title.Text = "Accept-Reject Outputs"
	h, err = NewHistogram(acceptData, 50)
	h.Width = 0.1
	h.Color = color.RGBA{R: 255, A: 255}

	if err != nil {
		log.Panic(err)
	}
	h.Normalize(1)
	plot.Add(h)

	// The normal distribution function
	plot.Add(norm)

	err = plot.Save(400, 400, "accept_reject_histogram.png")
	if err != nil {
		log.Panic(err)
	}

	// Generate the special properties numbers
	for i := 0; i < Executions; i++ {
		x := Polar(r)
		y := StdNorm(x)
		point := struct {
			X float64
			Y float64
		}{x, y}

		polarData = append(polarData, point)
	}

	// Generate the histogram
	plot, err = plotLib.New()
	if err != nil {
		log.Fatal(err)
	}
	plot.Title.Text = "Special Properties Outputs"
	h, err = NewHistogram(polarData, 50)
	h.Width = 0.1
	h.Color = color.RGBA{G: 255, A: 255}
	if err != nil {
		log.Panic(err)
	}
	h.Normalize(1)
	plot.Add(h)

	// The normal distribution function
	plot.Add(norm)

	err = plot.Save(400, 400, "polar_histogram.png")
	if err != nil {
		log.Panic(err)
	}

}
