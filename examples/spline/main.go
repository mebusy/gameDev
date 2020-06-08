package main

import (
    "github.com/mebusy/simpleui"
    "github.com/go-gl/glfw/v3.1/glfw"
    "image"
    "image/color"
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
    graph.FillRect( self.screenImage , self.screenImage.Bounds(), color.Black )

    // update control point 
    pt := spline_path.GetSelectedPoint()
    if pt != nil &&  simpleui.IsMouseKeyHold( window, glfw.MouseButtonLeft ) {
        mx,my := simpleui.GetCursorPosInWindow(window)
        pt.SetPosition( mx/float64(scale), my/float64(scale) )
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
    spline_path.Draw( self.screenImage, true, color.White )

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
}
func (self *MyView) TextureBuff() []uint8 {
    return self.screenImage.Pix
}
func (self *MyView) Title() string {
    return "my game"
}

var scale int
func main() {
    nPoint := 10
    points := make( []spline.Point2D, nPoint )
    var cx, cy float64 =  50, 40
    for i:=0; i<nPoint; i++ {
        // pt := Point2D{ float64(10+i*10), 41  }
        pt := spline.Point2D{
            X:cx+ 30*math.Sin(float64(i)/float64(nPoint)*2*math.Pi ),
            Y:cy+ 30*math.Cos(float64(i)/float64(nPoint)*2*math.Pi ) }
        points[i] = pt
    }
    spline_path = spline.NewSpline( points )

    var w,h int
    w,h,scale = 128,80,8
    view := NewView(w,h)
    simpleui.SetWindow( w,h, scale  )
    simpleui.Run( view )
}


var fMarker float64
var spline_path *spline.Spline
