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
    for i, pt := range self.points {
        if i == nSelectedPoint {
            pt.Draw( dst, COLOR_YELLOW )
        } else {
            pt.Draw( dst, color.White )
        }
    }
}


var spline_path Spline
var nSelectedPoint int

func init() {
    spline_path.points = []Point2D {  {10,41}, {40,41}, {70,41}, {100,41 } }

}

func switchControlPoint() {
    nSelectedPoint = (nSelectedPoint+1)% len(spline_path.points)
}


var (
    COLOR_YELLOW = color.RGBA{ 255,255,0,255 }
)

