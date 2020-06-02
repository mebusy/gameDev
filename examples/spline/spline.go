package main

import (
    "image/color"
    "image"
)

type Point2D struct {
    x,y float32
}

func (self *Point2D) Draw( dst *image.RGBA, c color.Color ) {
    w,h := dst.Bounds().Size().X , dst.Bounds().Size().Y
    x,y := int(self.x) , int(self.y)
    for j:=y-1; j<=y+1; j++ {
        for i:=x-1; i<=x+1; i++ {
            if i<0 || j<0 || i>=w || j>= h {
                continue
            }
            dst.Set( i,j, c )
        }
    }
}

type Spline struct {
    points []Point2D
}

func (self *Spline) Draw( dst *image.RGBA ) {
    for _, pt := range self.points {
        pt.Draw( dst, color.White )
    }
}


var spline_path Spline
var nSelectedPoint int

func init() {
    spline_path.points = []Point2D {  {10,41}, {40,41}, {70,41}, {100,41 } }

}



