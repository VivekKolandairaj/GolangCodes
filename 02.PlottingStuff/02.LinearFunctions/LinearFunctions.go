package main

import (
	"image/color"
	"log"
	
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	//create a new plot
	p := plot.New()

	//set the title and axis labels
	p.Title.Text = "Linear Functions"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	//set the x and y axis ranges
	//x axis range is [-5,5]
	p.X.Min = -5
	p.X.Max = 5
	//y axis range is [-11,11]
	p.Y.Min = -11
	p.Y.Max = 11

	//define the first function
	Line1Plot := plotter.NewFunction(line1)
	Line1Plot.Color = color.RGBA{R: 30, G: 150, B: 80, A: 255}
	Line1Plot.Width = vg.Points(2)
	Line1Plot.Samples = 500

	//define the second function
	Line2Plot := plotter.NewFunction(line2)
	Line2Plot.Color = color.RGBA{R: 140, G: 70, B: 190, A: 255}
	Line2Plot.Width = vg.Points(2)
	Line2Plot.Samples = 500

	//add the lines to the plot
	p.Add(Line1Plot, Line2Plot)

	//add the legend
	p.Legend.Add("y=2x+1", Line1Plot)
	p.Legend.Add("y=-x+3", Line2Plot)
	p.Legend.Top = false

	// save the plot as a png image
	if err := p.Save(8*vg.Inch, 5*vg.Inch, "Linear.png"); err != nil {
		log.Fatal(err)
	}
	log.Println("Plot saved to Linear.png")

}

// first function to be plotted
func line1(x float64) (y float64) {

	y = 2*x + 1
	return
}

// second function to be plotted
func line2(x float64) (y float64) {

	y = -x + 3
	return
}
