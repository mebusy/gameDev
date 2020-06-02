package main

import (
    "image"
    "image/color"
    "github.com/go-gl/glfw/v3.1/glfw"
    "github.com/mebusy/simpleui"
    // "log"
)

type Point2D struct {
    x,y float32
}

func (self *Point2D) Draw( dst *image.RGBA, c color.Color, radii int ) {
    w,h := dst.Bounds().Size().X , dst.Bounds().Size().Y
    x,y := int(self.x) , int(self.y)

    half_radii := radii/2
    for j:=y-half_radii; j<=y+half_radii; j++ {
        for i:=x-half_radii; i<=x+half_radii; i++ {
            if i<0 || j<0 || i>=w || j>= h {
                continue
            }
            dst.Set( i,j, c )
        }
    }
}
func (self *Point2D) Update( dt float64 ) {
    window := simpleui.GetWindow()
    distance := float32(30 * dt)
    // log.Println(dt)
    if simpleui.ReadKey(window, glfw.KeyLeft) {
        self.x -= distance
    }
    if simpleui.ReadKey(window, glfw.KeyRight) {
        self.x += distance
    }
    if simpleui.ReadKey(window, glfw.KeyUp) {
        self.y -= distance
    }
    if simpleui.ReadKey(window, glfw.KeyDown) {
        self.y += distance
    }


}

type Spline struct {
    points []Point2D
}



func (self *Spline) Update( dt float64 ) {
    self.points[nSelectedPoint].Update(dt)

    // update agent
    window := simpleui.GetWindow()
    if simpleui.ReadKey(window, glfw.KeyA) {
        fMarker -= float32(5 * dt)
    }
    if simpleui.ReadKey(window, glfw.KeyD) {
        fMarker += float32(5 * dt)
    }

    if fMarker >= float32( len(self.points) ) {
        fMarker -= float32( len(self.points) )
    }
    if fMarker < 0 {
        fMarker += float32( len(self.points) )
    }
}

func (self *Spline) Draw( dst *image.RGBA ) {
    // draw curve
    var t float32
    for t=0.0; t<float32(len(self.points)); t+= 0.01 {
        pt := self.getSplinePoint(t, true)
        pt.Draw( dst, color.White, 1 )
    }
    // Draw control point
    for i, pt := range self.points {
        if i == nSelectedPoint {
            pt.Draw( dst, COLOR_YELLOW, 3 )
        } else {
            pt.Draw( dst, COLOR_RED, 3 )
        }
    }
    // draw agent
    p1 := self.getSplinePoint( fMarker, true )
    s1 := self.getSplineSlope( fMarker, true )

    _ = p1
    _ = s1

}


var spline_path Spline
var nSelectedPoint int
var fMarker float32

func init() {
    spline_path.points = make( []Point2D, 0 )
    for i:=0; i<10; i++ {
        pt := Point2D{ float32(10+i*10), 41  }
        spline_path.points = append( spline_path.points , pt )
    }
}

func switchControlPoint() {
    nSelectedPoint = (nSelectedPoint+1)% len(spline_path.points)
}

var (
    COLOR_YELLOW = color.RGBA{ 255,255,0,255 }
    COLOR_RED = color.RGBA{ 255,0,0,255 }
)

