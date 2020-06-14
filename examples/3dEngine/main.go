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
    var matRotZ , matRotX m3d.Mat
    fTheta += dt
    matRotZ.Set( 0,0, math.Cos(fTheta)  )
    matRotZ.Set( 1,0, math.Sin(fTheta)  )
    matRotZ.Set( 0,1, -math.Sin(fTheta)  )
    matRotZ.Set( 1,1, math.Cos(fTheta)  )
    matRotZ.Set( 2,2, 1 )
    matRotZ.Set( 3,3, 1 )

    matRotX.Set( 0,0, 1 )
    matRotX.Set( 1,1, math.Cos(fTheta*0.5) )
    matRotX.Set( 2,1, math.Sin(fTheta*0.5) )
    matRotX.Set( 1,2, -math.Sin(fTheta*0.5) )
    matRotX.Set( 2,2, math.Cos(fTheta*0.5) )
    matRotX.Set( 3,3, 1 )

    // draw
    var triProj,triRotZ,triRotZX m3d.Triangle
    var tri2D = graph.NewTriangle(0,0,0,0,0,0)
    for _, tri := range meshCube.Tris {
        for i:=0;i<3;i++ {

            // rot z
            m3d.MultiplyMatrixVector( matRotZ, tri.P[i], &triRotZ.P[i] )
            // rot x
            m3d.MultiplyMatrixVector( matRotX, triRotZ.P[i], &triRotZX.P[i] )
            // debug , translate the trianagle + 3z
            triRotZX.P[i].Z += 3

            // projection
            m3d.MultiplyMatrixVector( matProj, triRotZX.P[i], &triProj.P[i] )

            // scale into view
            // 1. shift a coordinate to between [0,1]
            // 2. scale it to the appropriate size
            x2d := (triProj.P[i].X + 1) * 0.5 * float64(screenW)
            y2d := (triProj.P[i].Y + 1) * 0.5 * float64(screenH)
            tri2D.SetVert(i, int( x2d  ), int( y2d ) )
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
    screenW,screenH,scale = 256,128,2
    view := NewView(screenW,screenH)
    simpleui.SetWindow( screenW,screenH , scale  )
    simpleui.Run( view )
}

var meshCube m3d.Mesh
var matProj m3d.Mat
var fTheta float64
