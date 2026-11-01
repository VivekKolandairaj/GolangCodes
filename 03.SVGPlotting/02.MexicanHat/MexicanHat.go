// program for SVG rendering of the Mexican Hat (Ricker Wavelet) function
// f(x,y) = (1 - r^2) * e^(-r^2/2) where r = square root(x^2 + y^2)

// This is based on the program surface.go in pages 58 and 59 of 
// The Go Programming Language by Alan A.A. Donovan and Brian W. Kernighan

package main

// just import the fmt and math packages. No specialized plotting packages are necessary
import (
    "fmt"
    "math"
)

// constants required by the program
// the constants xyrange, xyscale, and zscale can be adjusted to get a better rendering of the function
const (
    width, height = 600, 320           // canvas size width: 600 px, height: 320 px
    cells         = 100                // number of grid cells (100 x 100 = 10,000 cells)
    xyrange       = 8.0               // axis ranges (-xyrange .. +xyrange)
    xyscale       = width / 2 /xyrange // number of pixels per x or y unit
    zscale        = height * 0.45       // pixels per z unit. Factor 0.45 is arbitrary and can be adjusted
    angle         = math.Pi / 6        // angle of x and y axes with the horizontal (=30 degrees) for isometric projection
)

// calculate the sine and cosine of 30 degrees
// this is needed for transforming (x,y,z) coordinates to (sx,sy) coordinates i.e. coordinate transformation
var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// main function
// program execution starts here
func main() {
    // opening tag
    fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
               "style='stroke: grey; fill: white; stroke-width: 0.7' "+
               "width='%d' height='%d'>\n", width,height)

    // nested for loops for generating the polygon tags
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            // calculate the 4 corners
            ax, ay := corner(i+1,j)
            bx, by := corner(i,j)
            cx, cy := corner(i,j+1)
            dx, dy := corner(i+1,j+1)

            // generate the svg polygon
            fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",ax,ay,bx,by,cx,cy,dx,dy)

        } // inner loop

    } // outer loop


    // closing tag
    fmt.Println("</svg>")
}

// corner function for coordinate transformations
// (i,j) -> (x,y,z) -> (sx,sy)
// The inputs are the two integers i and j
// The outpus are two float64 values sx and sy
func corner(i,j int) (float64,float64) {
    // find point (x,y) at corner of cell (i,j)
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    // compute the surface height z
    z := f(x,y)

    // project (x,y,z) isometrically onto the 2D SVG canvas  (sx,sy)
    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale

    // return the sx and sy values
    return sx,sy

}

// function f for computing z = f(x,y) = (1 - r^2) * e^(-r^2/2) where r = square root(x^2 + y^2)
// The inputs are two float64 values x and y
// The output is a float64 value z
func f(x,y float64) float64 {
    r := math.Hypot(x,y) // hypotenuse of x and y
    rSquared := math.Pow(r,2)
    return (1 - rSquared) * math.Exp(-rSquared/2) 

}

// go does not have a symbol for raising to a power
// use math.Pow()
// for exponential function use math.Exp()
