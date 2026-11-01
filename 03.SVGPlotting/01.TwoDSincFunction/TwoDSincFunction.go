// Plotting 2D sinc function
// From pages 58 and 59 of The Go Programming language book by 
// Alan A. A. Donovan and Brian W. Kernighan

// A dedicated package library is not used
// This program generates XML tags with coordinates for the corners of polygons
// The polygons create a wireframe model of the surface

package main

// only the fmt and math packages are imported
// The Printf function from fmt package prints the XML tags
// Sin, Cos, and Hypot functions are used from the math package

import (
    "fmt"
    "math"
)

//constants used in the program
const (
    width, height = 600, 320            // canvas size. Width: 600 px, height = 320 px
    cells         = 100                 // number of grid cells 100 x 100 = 10,000 cells
    xyrange       = 30.0                // axis ranges (-xyrange .. + xyrange)
    xyscale       = width / 2 / xyrange // pixels per x or y unit
    zscale        = height * 0.4        // pixels per z unit. Factor 0.4 is arbitrary
    angle         = math.Pi / 6         // angle of x and y axes (=30 degree) for isometric projection
)

// calculate the values of sine 30 degrees and cos 30 degrees here
// this is calculated once here and then used multiple times in the program

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// the main function
func main() {
    // opening tag
    fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' " +
               "style='stroke: grey; fill: white; stroke-width: 0.7' " +
               "width='%d' height='%d'>", width,height)
    fmt.Println() // newline

    // nested for loops for calculating the corners of the polygon
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            // calculate the four corners

            ax,ay := corner(i+1,j)
            bx,by := corner(i,j)
            cx,cy := corner(i,j+1)
            dx,dy := corner(i+1,j+1)

            fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",ax,ay,bx,by,cx,cy,dx,dy)
        } // inner loop
    } // outer loop

    // inner loop executes 100 times for every iteration of the outer loop
    // in total there are 100 x 100 = 10,000 iterations
    // each iteration processes one grid cell

    // closing tag
    fmt.Println("</svg>")
}

// function for calculating the corners of the polygon
// coordinate transform from (i,j) -> (x,y,z) -> (sx,sy)
func corner(i,j int) (float64,float64) {
    // find point (x,y) at corner of cell (i,j)
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    // compute surface height z
    z := f(x,y)

    // project (x,y,z) isometrically onto 2D SVG canvas (sx,sy)
    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale

    return sx,sy

}



// function for calculating z=f(x,y)=sin(r)/r where r = square root of the sum of x squared and y squared
// inputs are (x,y) coordinates and the output is the z coordinate
func f(x,y float64) float64 {
    r := math.Hypot(x,y) // hypotenuse i.e. distance from (0,0)

    if r == 0 {
        return 1
    } // this avoids the NaN issue when r is 0 

    return math.Sin(r)/r
}
