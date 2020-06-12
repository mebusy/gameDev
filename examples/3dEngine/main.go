package main

import (
    "github.com/mebusy/simpleui"
    "github.com/go-gl/glfw/v3.1/glfw"
    "image"
    // "image/color"
    "github.com/mebusy/simpleui/graph"
    // "image/draw"
    "math/rand"
    // "log"
    "time"
    "m3d"
)

type MyView struct {
    screenImage *image.RGBA
}

func NewView( w,h int) *MyView {
    view := &MyView{}
    view.screenImage = image.NewRGBA(image.Rect(0, 0, w, h))
    return view
}

var test_pts [6]int
func (self *MyView) Enter() {
    rand.Seed( time.Now().Unix() )

    meshCube.Tris = []m3d.Triangle {
        // SOUTH face . FRONT
        m3d.Triangle{ [3]m3d.Vec3D{{0.0, 0.0, 0.0},    {0.0, 1.0, 0.0},    {1.0, 1.0, 0.0} }} ,
        m3d.Triangle{ [3]m3d.Vec3D{{0.0, 0.0, 0.0},    {1.0, 1.0, 0.0},    {1.0, 0.0, 0.0} }},
        // EAST
        m3d.Triangle{ [3]m3d.Vec3D{{1.0, 0.0, 0.0},    {1.0, 1.0, 0.0},    {1.0, 1.0, 1.0} }},
        m3d.Triangle{ [3]m3d.Vec3D{{1.0, 0.0, 0.0},    {1.0, 1.0, 1.0},    {1.0, 0.0, 1.0} }},
        // NORTH
        m3d.Triangle{ [3]m3d.Vec3D{{1.0, 0.0, 1.0},    {1.0, 1.0, 1.0},    {0.0, 1.0, 1.0} }},
        m3d.Triangle{ [3]m3d.Vec3D{{1.0, 0.0, 1.0},    {0.0, 1.0, 1.0},    {0.0, 0.0, 1.0} }},
        // EAST
        m3d.Triangle{ [3]m3d.Vec3D{{0.0, 0.0, 1.0},    {0.0, 1.0, 1.0},    {0.0, 1.0, 0.0} }},
        m3d.Triangle{ [3]m3d.Vec3D{{0.0, 0.0, 1.0},    {0.0, 1.0, 0.0},    {0.0, 0.0, 0.0} }},
        // TOP
        m3d.Triangle{ [3]m3d.Vec3D{{0.0, 1.0, 0.0},    {0.0, 1.0, 1.0},    {1.0, 1.0, 1.0} }},
        m3d.Triangle{ [3]m3d.Vec3D{{0.0, 1.0, 0.0},    {1.0, 1.0, 1.0},    {1.0, 1.0, 0.0} }},
        // BOTTOM
        m3d.Triangle{ [3]m3d.Vec3D{{1.0, 0.0, 1.0},    {0.0, 0.0, 1.0},    {0.0, 0.0, 0.0} }},
        m3d.Triangle{ [3]m3d.Vec3D{{1.0, 0.0, 1.0},    {0.0, 0.0, 0.0},    {1.0, 0.0, 0.0} }},
    }
}
func (self *MyView) Exit() {}
func (self *MyView) Update(t, dt float64) {

    graph.FillRect( self.screenImage, self.screenImage.Bounds() ,
                graph.COLOR_BLACK )


}

func (self *MyView) OnKey(key glfw.Key) {}
func (self *MyView) TextureBuff() []uint8 {
    return self.screenImage.Pix
}
func (self *MyView) Title() string {
    return "3d engine"
}


var screenW, screenH, scale int

func main() {
    screenW,screenH,scale = 256,240,2
    view := NewView(screenW,screenH)
    simpleui.SetWindow( screenW,screenH , scale  )
    simpleui.Run( view )
}

var meshCube m3d.Mesh

