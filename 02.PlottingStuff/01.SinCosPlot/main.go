//Go program for plotting sine and cosine curves

package main

import (
	"image/color"
	"log"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	//create a new plot
	p := plot.New()

	//set the title and axis labels
	p.Title.Text = "Sine and Cosine functions"
	p.X.Label.Text = "x (radians)"
	p.Y.Label.Text = "y"

	//set the visible range
	p.X.Min = -2 * math.Pi
	p.X.Max = 2 * math.Pi
	p.Y.Min = -1.2
	p.Y.Max = 1.2

	//Define the sine function
	sinPlot := plotter.NewFunction(func(x float64) float64 { return math.Sin(x) })
	sinPlot.Color = color.RGBA{R: 220, G: 50, B: 50, A: 255}
	sinPlot.Width = vg.Points(1.5)
	sinPlot.Samples = 500

	//Define the cosine function
	cosPlot := plotter.NewFunction(func(x float64) float64 { return math.Cos(x) })
	cosPlot.Color = color.RGBA{R: 30, G: 100, B: 220, A: 255}
	cosPlot.Width = vg.Points(1.5)
	cosPlot.Samples = 500

	//Add the curves to the plot
	p.Add(sinPlot, cosPlot)

	//Add the legend
	p.Legend.Add("sin(x)", sinPlot)
	p.Legend.Add("cos(x)", cosPlot)
	p.Legend.Top = true

	//save the plot as a png image
	if err := p.Save(8*vg.Inch, 5*vg.Inch, "SinCos.png"); err != nil {
		log.Fatal(err)
	}

	log.Println("Plot saved to SinCos.png")

}
