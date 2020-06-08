package main

import (
	"image"
	"image/color"

	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/mebusy/simpleui"

	// "image/color"
	// "image/draw"
	"math"
	"spline"

	"github.com/mebusy/simpleui/graph"
)

type MyView struct {
    screenImage *image.RGBA
}

func NewView( w,h int) *MyView {
    view := &MyView{}
    view.screenImage = image.NewRGBA(image.Rect(0, 0, w, h))
    return view
}

func (self *MyView) Enter() {
    window := simpleui.GetWindow()
    window.SetMouseButtonCallback( func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {

            if action == glfw.Press && button == glfw.MouseButtonLeft {
                mx,my := simpleui.GetCursorPosInWindow(window)
                spline_path.SelectControlPoint( mx/float64(scale),my/float64(scale) )
            } else if action == glfw.Release && button == glfw.MouseButtonLeft {
                spline_path.SelectControlPoint( -10, -10)
            }
        },
    )

}
func (self *MyView) Exit() {}
func (self *MyView) Update( gt, dt float64) {
    window := simpleui.GetWindow()
    // clear
    graph.FillRect( self.screenImage , self.screenImage.Bounds(), graph.COLOR_GREEN)

    // update control point 
    pt := spline_path.GetSelectedPoint()
    if pt != nil &&  simpleui.IsMouseKeyHold( window, glfw.MouseButtonLeft ) {
        mx,my := simpleui.GetCursorPosInWindow(window)
        pt.SetPosition( mx/float64(scale), my/float64(scale) )
    }

    // Calculate track boundary points
    fTrackWidth := 10.0
    for i:=0; i< len(spline_path.Ctl_points); i++ {
        p1 := spline_path.GetSplinePoint( float64(i), true )
        s1 := spline_path.GetSplineSlope( float64(i), true )
        slen := math.Sqrt( s1.X*s1.X + s1.Y*s1.Y )

        spline_trackleft.Ctl_points[i].X = p1.X + fTrackWidth* (-s1.Y/slen)
        spline_trackleft.Ctl_points[i].Y = p1.Y + fTrackWidth* ( s1.X/slen)
        spline_trackright.Ctl_points[i].X = p1.X - fTrackWidth* (-s1.Y/slen)
        spline_trackright.Ctl_points[i].Y = p1.Y - fTrackWidth* ( s1.X/slen)
    }

    // Reset racing line
    for i:=0; i< len(spline_racingline.Ctl_points); i++ {
        spline_racingline.Ctl_points[i] = spline_path.Ctl_points[i]
        displacement[i] = 0
    }

    for n:=0; n< nIterations; n++ {
        for i:=0; i<len(spline_racingline.Ctl_points); i++ {

        }
        // clamp displaced points to track width
        for i:=0; i<len(spline_racingline.Ctl_points); i++ {
            if displacement[i] > fTrackWidth {
                displacement[i] = fTrackWidth
            }
            if displacement[i] < -fTrackWidth {
                displacement[i] = -fTrackWidth
            }
            s := spline_path.GetSplineSlope( float64(i), true )
            slen := math.Sqrt( s.X*s.X + s.Y*s.Y )
            s.X /= slen ;  s.Y /= slen

            spline_racingline.Ctl_points[i].X = spline_path.Ctl_points[i].X +
                (-s.Y)*displacement[i]
            spline_racingline.Ctl_points[i].Y = spline_path.Ctl_points[i].Y +
                ( s.X)*displacement[i]
        }
    }

    // update agent
    if simpleui.ReadKey(window, glfw.KeyA) {
        nIterations ++
    }
    if simpleui.ReadKey(window, glfw.KeyD) {
        nIterations --
        if nIterations < 0 {
            nIterations = 0
        }
    }

    if fMarker >= spline_path.TotalSplineLength {
        fMarker -= spline_path.TotalSplineLength
    }
    if fMarker < 0 {
        fMarker += spline_path.TotalSplineLength
    }

    // draw track 
    fRes := 0.2
    var t float64
    var triangle = graph.NewTriangle(0,0,0,0,0,0)
    for t=0; t< float64(len( spline_path.Ctl_points)) ; t+= fRes {
        pl1 := spline_trackleft.GetSplinePoint( t, true )
        pr1 := spline_trackright.GetSplinePoint( t, true )
        pl2 := spline_trackleft.GetSplinePoint( t+fRes, true )
        pr2 := spline_trackright.GetSplinePoint( t+fRes, true )

        triangle.SetVert( 0, int(pl1.X ), int(pl1.Y)  )
        triangle.SetVert( 1, int(pr1.X ), int(pr1.Y)  )
        triangle.SetVert( 2, int(pr2.X ), int(pr2.Y)  )
        graph.FillTriangle( self.screenImage, triangle, graph.COLOR_GRAY )

        triangle.SetVert( 0, int(pl1.X ), int(pl1.Y)  )
        triangle.SetVert( 1, int(pl2.X ), int(pl2.Y)  )
        triangle.SetVert( 2, int(pr2.X ), int(pr2.Y)  )
        graph.FillTriangle( self.screenImage, triangle, graph.COLOR_GRAY )
    }

    // draw spline line
    spline_path.Draw( self.screenImage, true, color.White )
    // spline_trackleft.Draw( self.screenImage, false )
    // spline_trackright.Draw( self.screenImage, false )
    spline_racingline.Draw( self.screenImage, false, graph.COLOR_BLUE )



    // draw agent
    offset := spline_path.GetNormalizedOffset( fMarker )
    p1 := spline_path.GetSplinePoint( offset, true )
    s1 := spline_path.GetSplineSlope( offset, true )
    // orthogonal 
    r := math.Atan2( float64(-s1.X), float64(s1.Y) )
    nLen := 5.0
    graph.DrawLine(self.screenImage,
        int(float64(nLen)*math.Cos(r)+p1.X), int(float64(nLen)*math.Sin(r)+p1.Y),
        int(-float64(nLen)*math.Cos(r)+p1.X), int(-float64(nLen)*math.Sin(r)+p1.Y),
        graph.COLOR_BLUE)
}

func (self *MyView) SetAudioDevice(audio *simpleui.Audio) {}
func (self *MyView) OnKey(key glfw.Key) {
    // switch key {
    // case glfw.KeyX:
    //     spline_path.SwitchControlPoint()
    // }
}
func (self *MyView) TextureBuff() []uint8 {
    return self.screenImage.Pix
}
func (self *MyView) Title() string {
    return "my game"
}

var scale int
func main() {
    var w,h int
    w,h,scale = 256,240,4
    view := NewView(w,h)
    simpleui.SetWindow( w,h, scale  )

    for t:=0; t<4; t++ {
        nPoint := 20
        points := make( []spline.Point2D, nPoint )
        cx, cy :=  w/2, h/2
        for i:=0; i<nPoint; i++ {
            // pt := Point2D{ float64(10+i*10), 41  }
            pt := spline.Point2D{
                X:float64(cx)+ 80*math.Sin(float64(i)/float64(nPoint)*2*math.Pi ),
                Y:float64(cy)+ 80*math.Cos(float64(i)/float64(nPoint)*2*math.Pi ) }
            points[i] = pt
        }

        switch t {
        case 0:
            spline_path = spline.NewSpline( points )
        case 1:
            spline_trackleft = spline.NewSpline( points )
        case 2:
            spline_trackright = spline.NewSpline( points )
        case 3:
            spline_racingline = spline.NewSpline( points )
        }
    }


    simpleui.Run( view )
}


var fMarker float64
var spline_path *spline.Spline
var spline_trackleft *spline.Spline
var spline_trackright *spline.Spline
var spline_racingline *spline.Spline
var displacement [20]float64
var nIterations int = 1



