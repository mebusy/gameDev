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

// ===================================

type Spline struct {
    ctl_points []Point2D
    ctl_pt_lengths []float64
    TotalSplineLength float64
    nSelectedPoint int
}

func NewSpline( points []Point2D  ) *Spline {
    spl := &Spline{}

    nPoint := len(points)
    spl.ctl_points = make( []Point2D, nPoint )
    spl.ctl_pt_lengths = make( []float64, nPoint )

    for i:=0; i<nPoint; i++ {
        spl.ctl_points[i] = points[i]
    }
    log.Println( "new spline", nPoint)
    return spl
}

func (self *Spline) GetSelectedPoint()  *Point2D {
    return &self.ctl_points[self.nSelectedPoint]
}

func (self *Spline) SelectControlPoint( mx,my float64 )  *Point2D {
    for i,pt := range self.ctl_points {
        if math.Abs( pt.X - mx ) < 5 && math.Abs( pt.Y - my ) < 5 {
            self.nSelectedPoint = i
            return &pt
        }
    }
    return nil
}

func (self *Spline) SwitchControlPoint() {
    self.nSelectedPoint = (self.nSelectedPoint+1)% len(self.ctl_points)
}

func (self *Spline) Draw( dst *image.RGBA ) {
    // draw curve
    var t float64
    for t=0.0; t<float64(len(self.ctl_points)); t+= 0.01 {
        pt := self.GetSplinePoint(t, true)
        pt.Draw( dst, color.White, 1 )
    }
    // Draw control point
    self.TotalSplineLength = 0
    for i, pt := range self.ctl_points {
        self.ctl_pt_lengths[i] = self.CalculateSegmentLength( i, true )
        self.TotalSplineLength += self.ctl_pt_lengths[i]
        if i == self.nSelectedPoint {
            pt.Draw( dst, graph.COLOR_YELLOW, 3 )
        } else {
            pt.Draw( dst, graph.COLOR_RED, 3 )
        }
    }

}





