package spline

import (
    "image"
    "image/color"
    "math"
    "github.com/mebusy/simpleui/graph"
    "log"
)


// ===================================

type Point2D struct {
    X,Y float64
}

func (self *Point2D) Draw( dst *image.RGBA, c color.Color, radii int ) {
    w,h := dst.Bounds().Size().X , dst.Bounds().Size().Y
    x,y := int(self.X) , int(self.Y)

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

func (self *Point2D) DistanceTo( pt Point2D ) float64 {
    dx := (self.X - pt.X)
    dy := (self.Y - pt.Y)
    return math.Sqrt( float64(dx*dx + dy*dy) )
}

func (self *Point2D) SetPosition( x, y float64 ) {
    self.X = x
    self.Y = y
}
func (self Point2D) Normalized( ) Point2D {
    vecLen := math.Sqrt(  self.X*self.X + self.Y * self.Y )
    return Point2D{ self.X/vecLen, self.Y/vecLen }
}

// ===================================

type Spline struct {
    Ctl_points []Point2D
    ctl_pt_lengths []float64
    TotalSplineLength float64
    nSelectedPoint int
    Size int
}

func NewSpline( points []Point2D  ) *Spline {
    spl := &Spline{}

    nPoint := len(points)
    spl.Ctl_points = make( []Point2D, nPoint )
    spl.ctl_pt_lengths = make( []float64, nPoint )
    spl.nSelectedPoint = -1
    spl.Size = nPoint

    for i:=0; i<nPoint; i++ {
        spl.Ctl_points[i] = points[i]
    }
    log.Println( "new spline", nPoint)
    return spl
}

func (self *Spline) GetSelectedPoint()  *Point2D {
    if self.nSelectedPoint < 0 {
        return nil
    }
    return &self.Ctl_points[self.nSelectedPoint]
}

func (self *Spline) SelectControlPoint( mx,my float64 )  *Point2D {
    for i,pt := range self.Ctl_points {
        if math.Abs( pt.X - mx ) < 5 && math.Abs( pt.Y - my ) < 5 {
            self.nSelectedPoint = i
            return &pt
        }
    }
    self.nSelectedPoint = -1
    return nil
}


func (self *Spline) Draw( dst *image.RGBA, bDrawCtlPoints bool, line_color color.Color ) {
    // draw curve
    var t float64
    for t=0.0; t<float64(len(self.Ctl_points)); t+= 0.01 {
        pt := self.GetSplinePoint(t, true)
        pt.Draw( dst, line_color , 1 )
    }
    // Draw control point
    self.TotalSplineLength = 0
    for i, pt := range self.Ctl_points {
        self.ctl_pt_lengths[i] = self.CalculateSegmentLength( i, true )
        self.TotalSplineLength += self.ctl_pt_lengths[i]

        if bDrawCtlPoints {
            if i == self.nSelectedPoint {
                pt.Draw( dst, graph.COLOR_YELLOW, 3 )
            } else {
                pt.Draw( dst, graph.COLOR_RED, 3 )
            }
        }
    }

}





