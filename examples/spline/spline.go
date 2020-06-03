package main

import (
	"image"
	"image/color"
	"math"

	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/mebusy/simpleui"
	"github.com/mebusy/simpleui/graph"
	// "log"
)


var spline_path Spline
var nSelectedPoint int
var fMarker float32

func init() {
    nPoint := 10
    spline_path.ctl_points = make( []Point2D, nPoint )
    spline_path.ctl_pt_lengths = make( []float32, nPoint )
    var cx, cy float32 =  50, 40
    for i:=0; i<nPoint; i++ {
        // pt := Point2D{ float32(10+i*10), 41  }
        pt := Point2D{ 
            cx+ float32(30*math.Sin(float64(i)/float64(nPoint)*2*math.Pi )),
            cy+ float32(30*math.Cos(float64(i)/float64(nPoint)*2*math.Pi )) }
        spline_path.ctl_points[i] = pt
    }
}

// ===================================

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

func (self *Point2D) DistanceTo( pt Point2D ) float32 {
    dx := (self.x - pt.x)
    dy := (self.y - pt.y)
    return float32(math.Sqrt( float64(dx*dx + dy*dy) ))
}

// ===================================

type Spline struct {
    ctl_points []Point2D
    ctl_pt_lengths []float32
    totalSplineLength float32
}



func (self *Spline) Update( dt float64 ) {
    self.ctl_points[nSelectedPoint].Update(dt)

    // update agent
    window := simpleui.GetWindow()
    if simpleui.ReadKey(window, glfw.KeyA) {
        fMarker -= float32(20 * dt)
    }
    if simpleui.ReadKey(window, glfw.KeyD) {
        fMarker += float32(20 * dt)
    }

    if fMarker >= self.totalSplineLength {
        fMarker -= self.totalSplineLength
    }
    if fMarker < 0 {
        fMarker += self.totalSplineLength
    }
}

func (self *Spline) Draw( dst *image.RGBA ) {
    // draw curve
    var t float32
    for t=0.0; t<float32(len(self.ctl_points)); t+= 0.01 {
        pt := self.getSplinePoint(t, true)
        pt.Draw( dst, color.White, 1 )
    }
    // Draw control point
    self.totalSplineLength = 0
    for i, pt := range self.ctl_points {
        self.ctl_pt_lengths[i] = self.CalculateSegmentLength( i, true )
        self.totalSplineLength += self.ctl_pt_lengths[i]
        if i == nSelectedPoint {
            pt.Draw( dst, graph.COLOR_YELLOW, 3 )
        } else {
            pt.Draw( dst, graph.COLOR_RED, 3 )
        }
    }
    // draw agent
    offset := self.GetNormalizedOffset( fMarker )
    p1 := self.getSplinePoint( offset, true )
    s1 := self.getSplineSlope( offset, true )
    // orthogonal 
    r := math.Atan2( float64(-s1.x), float64(s1.y) )
    nLen := 5.0
    graph.DrawLine(dst,
        nLen*math.Cos(r)+float64(p1.x), nLen*math.Sin(r)+float64(p1.y),
        -nLen*math.Cos(r)+float64(p1.x), -nLen*math.Sin(r)+float64(p1.y),
        graph.COLOR_BLUE)

}



func switchControlPoint() {
    nSelectedPoint = (nSelectedPoint+1)% len(spline_path.ctl_points)
}


