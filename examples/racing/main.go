package main

import (
    "github.com/mebusy/simpleui"
    "github.com/go-gl/glfw/v3.1/glfw"
    "image"
    // "image/color"
    // "image/draw"
    "github.com/mebusy/simpleui/graph"
    "spline"
    "math"
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
func (self *MyView) Update(t, dt float64) {
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


    // update agent
    if simpleui.ReadKey(window, glfw.KeyA) {
        fMarker -= 20 * dt
    }
    if simpleui.ReadKey(window, glfw.KeyD) {
        fMarker += 20 * dt
    }

    if fMarker >= spline_path.TotalSplineLength {
        fMarker -= spline_path.TotalSplineLength
    }
    if fMarker < 0 {
        fMarker += spline_path.TotalSplineLength
    }

    // draw spline line
    spline_path.Draw( self.screenImage, true )
    spline_trackleft.Draw( self.screenImage, false )
    spline_trackright.Draw( self.screenImage, false )

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

    for t:=0; t<3; t++ {
        nPoint := 20
        points := make( []spline.Point2D, nPoint )
        cx, cy :=  w/2, h/2
        for i:=0; i<nPoint; i++ {
            // pt := Point2D{ float64(10+i*10), 41  }
            pt := spline.Point2D{
                X:float64(cx)+ 30*math.Sin(float64(i)/float64(nPoint)*2*math.Pi ),
                Y:float64(cy)+ 30*math.Cos(float64(i)/float64(nPoint)*2*math.Pi ) }
            points[i] = pt
        }

        switch t {
        case 0:
            spline_path = spline.NewSpline( points )
        case 1:
            spline_trackleft = spline.NewSpline( points )
        case 2:
            spline_trackright = spline.NewSpline( points )
        }
    }


    simpleui.Run( view )
}


var fMarker float64
var spline_path *spline.Spline
var spline_trackleft *spline.Spline
var spline_trackright *spline.Spline
