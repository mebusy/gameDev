package main

import (
	"image"
	"math"

	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/mebusy/simpleui"

	// "image/color"
	"github.com/mebusy/simpleui/graph"
	// "image/draw"
	"math/rand"
	// "log"
	"m3d"
	"time"
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

    fZNear := 0.1
    fZFar := 1000.0
    fFov := 90.0 // degree
    fAspectRatio := float64(screenH)/float64(screenW)
    fFovRad := 1/math.Tan(  fFov*0.5 /180 *math.Pi  )

    matProj.Clear()
    matProj.Set(0,0, fAspectRatio * fFovRad )
    matProj.Set(1,1, fFovRad)
    matProj.Set(2,2, fZFar / (fZFar-fZNear))
    matProj.Set(2,3,-fZNear * fZFar / (fZFar-fZNear))
    matProj.Set(3,2, 1)
    matProj.Set(3,3, 0)
}
func (self *MyView) Exit() {}
func (self *MyView) Update(t, dt float64) {

    graph.FillRect( self.screenImage, self.screenImage.Bounds() ,
                graph.COLOR_BLACK )

    var triProj m3d.Triangle
    var tri2D = graph.NewTriangle(0,0,0,0,0,0)
    for _, tri := range meshCube.Tris {
        for i:=0;i<3;i++ {
            m3d.MultiplyMatrixVector( matProj, tri.P[i], &triProj.P[i] )
            tri2D.SetVert(i, int( tri.P[i].X ), int( tri.P[i].Y) )
        }
        graph.DrawTriangle( self.screenImage, tri2D, graph.COLOR_WHITE )
    }
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
var matProj m3d.Mat
