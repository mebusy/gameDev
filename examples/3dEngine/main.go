package main

import (
    "sort"
	"image"
	"log"

	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/mebusy/simpleui"

	"image/color"
	"github.com/mebusy/simpleui/graph"
	// "image/draw"
	"math/rand"
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

    meshCube.LoadFromObj( "axis2.obj"  )
    /*
    meshCube.Tris = []m3d.Triangle {
        // SOUTH face . FRONT
        m3d.Triangle{ [3]m3d.Vec3D{{0.0, 0.0, 0.0,1},    {0.0, 0.5, 0.0,1},    {1.0, 1.0, 0.0,1} },1} ,
        m3d.Triangle{ [3]m3d.Vec3D{{0.0, 0.0, 0.0,1},    {1.0, 1.0, 0.0,1},    {1.0, 0.0, 0.0,1} },1},
        // EAST
        m3d.Triangle{ [3]m3d.Vec3D{{1.0, 0.0, 0.0,1},    {1.0, 1.0, 0.0,1},    {1.0, 1.0, 1.0,1} },1},
        m3d.Triangle{ [3]m3d.Vec3D{{1.0, 0.0, 0.0,1},    {1.0, 1.0, 1.0,1},    {1.0, 0.0, 1.0,1} },1},
        // NORTH
        m3d.Triangle{ [3]m3d.Vec3D{{1.0, 0.0, 1.0,1},    {1.0, 1.0, 1.0,1},    {0.0, 1.0, 1.0,1} },1},
        m3d.Triangle{ [3]m3d.Vec3D{{1.0, 0.0, 1.0,1},    {0.0, 1.0, 1.0,1},    {0.0, 0.0, 1.0,1} },1},
        // EAST
        m3d.Triangle{ [3]m3d.Vec3D{{0.0, 0.0, 1.0,1},    {0.0, 1.0, 1.0,1},    {0.0, 0.5, 0.0,1} },1},
        m3d.Triangle{ [3]m3d.Vec3D{{0.0, 0.0, 1.0,1},    {0.0, 0.5, 0.0,1},    {0.0, 0.0, 0.0,1} },1},
        // TOP
        m3d.Triangle{ [3]m3d.Vec3D{{0.0, 0.5, 0.0,1},    {0.0, 1.0, 1.0,1},    {1.0, 1.0, 1.0,1} },1},
        m3d.Triangle{ [3]m3d.Vec3D{{0.0, 0.5, 0.0,1},    {1.0, 1.0, 1.0,1},    {1.0, 1.0, 0.0,1} },1},
        // BOTTOM
        m3d.Triangle{ [3]m3d.Vec3D{{1.0, 0.0, 1.0,1},    {0.0, 0.0, 1.0,1},    {0.0, 0.0, 0.0,1} },1},
        m3d.Triangle{ [3]m3d.Vec3D{{1.0, 0.0, 1.0,1},    {0.0, 0.0, 0.0,1},    {1.0, 0.0, 0.0,1} },1},
    }
    //*/

    // for OpenGL style, glFrustum accept only positive value of near and far
    // we need to negate them during the construction of GL_PROJECTION matrix
    fZNear := 5.0
    fZFar := 1000.0
    fFov := 60.0 // degree
    fAspectRatio := float64(screenW)/float64(screenH)

    matProj  = m3d.NewProjectionMat( fFov, fAspectRatio, fZNear, fZFar  )

    for _, vec := range (   []m3d.Vec3D{ {1,1,-2,1}, {1,1,-1000,1}, {1,1,2,1}, {1,1,1000,1} } ) {
        log.Printf( "project %v -> %v" , vec , m3d.MultiplyMatrixVector( matProj, vec )  )
    }

    // viewer matrix test
    mat_viewer := m3d.QuickInverse(m3d.NewCameraMat( m3d.Vec3D{ 1,2,3,1 } , m3d.Vec3D{ 4,3,4,1 }  , m3d.Vec3D{ 0,1,0,1 }   ))
    log.Printf( "test viewer: %+v " , mat_viewer    )

    log.Printf( "view*[11,2,6,1]= %+v", m3d.MultiplyMatrixVector( mat_viewer , m3d.Vec3D{ 11,2,6,1 }  )  )
    var ma, mb m3d.Mat
    ma.M = [16]float64{  1,0,0,0, 2,1,0,0, 3,4,5,6, 0,1,0,1 }
    mb.M = [16]float64{  1,2,3,4, 1,1,0,0, 7,6,5,6, 1,1,0,1 }
    log.Printf( "ma*mb, sould be 14 18 15 22 ... 3 2 0 1 %+v " , m3d.MultiplyMatrixMatrix( ma,mb )  )
}
func (self *MyView) Exit() {}
func (self *MyView) Update(t, dt float64) {
    windows := simpleui.GetWindow()
    if simpleui.ReadKey(windows, glfw.KeyUp) {
        vCamera.Y += 8 * dt
    }
    if simpleui.ReadKey(windows, glfw.KeyDown) {
        vCamera.Y -= 8 * dt
    }
    if simpleui.ReadKey(windows, glfw.KeyLeft) {
        vCamera.X -= 8 * dt
    }
    if simpleui.ReadKey(windows, glfw.KeyRight) {
        vCamera.X += 8 * dt
    }
    vForward := vLookDir.Mul( 8*dt )
    if simpleui.ReadKey(windows, glfw.KeyW) {
        vCamera = vCamera.Add( vForward )
    }
    if simpleui.ReadKey(windows, glfw.KeyS) {
        vCamera = vCamera.Sub( vForward )
    }

    if simpleui.ReadKey(windows, glfw.KeyA) {
        fYaw -= 2 * dt
    }
    if simpleui.ReadKey(windows, glfw.KeyD) {
        fYaw += 2 * dt
    }

    graph.FillRect( self.screenImage, self.screenImage.Bounds() ,
                graph.COLOR_BLACK )
    // fTheta += dt
    matRotZ := m3d.NewRotZMat( fTheta )
    matRotX := m3d.NewRotXMat( fTheta * 0.5 )

    matTrans := m3d.NewTransMat( 0,0, -5 )
    matWorld := m3d.NewIdentityMat()
    matWorld = m3d.MultiplyMatrixMatrix( matRotX , matRotZ )
    matWorld = m3d.MultiplyMatrixMatrix( matTrans , matWorld )

    /*
    vLookDir = m3d.Vec3D{ 0,0,1,1 }
    vUp := m3d.Vec3D{ 0,1,0,1 }
    vTarget := vCamera.Add(vLookDir)
    /*/
    vUp := m3d.Vec3D{ 0,1,0,1 }
    // for OpenGL style, camera is looking at -Z 
    vTarget := m3d.Vec3D{ 0,0,-1,1 }
    matCameraRot := m3d.NewRotYMat( fYaw )
    vLookDir = m3d.MultiplyMatrixVector( matCameraRot , vTarget )
    // log.Println(vLookDir)
    vTarget = vCamera.Add( vLookDir )
    //*/

    matCamera := m3d.NewCameraMat( vCamera, vTarget, vUp  )
    matView := m3d.QuickInverse( matCamera )
    // log.Printf( "matCamera:%+v", matCamera  )
    // log.Printf( "matView:%+v", matView  )
    // debug
    // matView = m3d.NewIdentityMat()

    triangles2Raster := make( []m3d.Triangle, 0 )

    // draw triangels
    var triTransformed, triViewed m3d.Triangle
    for _, tri := range meshCube.Tris {
        // rotate and transform
        for i:=0;i<3;i++ {
            triTransformed.P[i] = m3d.MultiplyMatrixVector( matWorld, tri.P[i] )
        }

        // calculate normal
        // draw only when triangle is visible
        normal := triTransformed.CalculateNormal()
        /*
        if normal.Z < 0 {
        /*/
        if normal.Dot( triTransformed.P[0].To(vCamera) ) > 0  {
        //*/
            // illumination 
            dp := normal.Dot( light_direction_normalized )
            // Convert World Space --> View Space
            for i:=0;i<3;i++ {
                triViewed.P[i] = m3d.MultiplyMatrixVector( matView, triTransformed.P[i] )
            }


            var triProj m3d.Triangle
            triProj.Color = dp
            visible := true
            for i:=0;i<3;i++ {
                // projection
                triProj.P[i] = m3d.MultiplyMatrixVector( matProj, triViewed.P[i] ) //.NormalizeByW()
                // just test
                if triProj.P[i].W < 0 {
                    visible = false
                    // break
                }
                triProj.P[i] = triProj.P[i].NormalizeByW()
            }
            if !visible {
                // continue
            }
            triangles2Raster = append( triangles2Raster , triProj )


        } // draw visible triangle
    } // visit triangle

    sort.SliceStable( triangles2Raster , func(i, j int) bool {
        return triangles2Raster[i].MidPointZ() > triangles2Raster[j].MidPointZ()
    })

    var tri2D = graph.NewTriangle(0,0,0,0,0,0)
    for _, triProj := range triangles2Raster {
        for i:=0;i<3;i++ {
            // scale into view
            // 1. shift a coordinate to between [0,1]
            // 2. scale it to the appropriate size
            x2d := (triProj.P[i].X + 1) * 0.5 * float64(screenW)
            y2d := (-triProj.P[i].Y + 1) * 0.5 * float64(screenH)
            tri2D.SetVert(i, int( x2d  ), int( y2d ) )
        }

        gray := uint8(triProj.Color*128 + 100  )
        graph.FillTriangle( self.screenImage, tri2D, color.RGBA{ gray,gray,gray,255 } )
        graph.DrawTriangle( self.screenImage, tri2D, graph.COLOR_RED )
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
    screenW,screenH,scale = 512,256,2
    view := NewView(screenW,screenH)
    simpleui.SetWindow( screenW,screenH , scale  )
    simpleui.Run( view )
}

var meshCube m3d.Mesh
var matProj m3d.Mat
var fTheta float64

var vCamera m3d.Vec3D = m3d.Vec3D{ X:0,Y:0,Z:0,W:1 }
var vLookDir m3d.Vec3D
var light_direction_normalized =  m3d.Vec3D{ X:0,Y:0,Z:-1,W:1 }

var fYaw float64



