package m3d

import (
    "math"
    "log"
)

func MultiplyMatrixVector( m Mat, i Vec3D ) Vec3D {
    // debug
    if i.W != 1 {
        println( "[debug]: vector with w:0" )
    }
    var o Vec3D
    o.X = i.X*m.M[(0<<2)+0]+i.Y*m.M[(0<<2)+1]+i.Z*m.M[(0<<2)+2]+i.W*m.M[(0<<2)+3]
    o.Y = i.X*m.M[(1<<2)+0]+i.Y*m.M[(1<<2)+1]+i.Z*m.M[(1<<2)+2]+i.W*m.M[(1<<2)+3]
    o.Z = i.X*m.M[(2<<2)+0]+i.Y*m.M[(2<<2)+1]+i.Z*m.M[(2<<2)+2]+i.W*m.M[(2<<2)+3]
    o.W = i.X*m.M[(3<<2)+0]+i.Y*m.M[(3<<2)+1]+i.Z*m.M[(3<<2)+2]+i.W*m.M[(3<<2)+3]

    /*
    // if w == 0, this point makes no sense
    // this is internally done by openGL
    if w != 0.0 {
        o.X /= w
        o.Y /= w
        o.Z /= w
    }
    //*/
    return o
}

func MultiplyMatrixMatrix( m1 Mat, m2 Mat ) Mat {
    var m Mat
    dest := &m.M  // use tmp variable for performance
    src1 := m1.M
    src2 := m2.M
    dest[(0<<2)+0] = src1[(0<<2)+0] * src2[(0<<2)+0] + src1[(0<<2)+1] * src2[(1<<2)+0] + src1[(0<<2)+2] * src2[(2<<2)+0] + src1[(0<<2)+3] * src2[(3<<2)+0] 
    dest[(0<<2)+1] = src1[(0<<2)+0] * src2[(0<<2)+1] + src1[(0<<2)+1] * src2[(1<<2)+1] + src1[(0<<2)+2] * src2[(2<<2)+1] + src1[(0<<2)+3] * src2[(3<<2)+1] 
    dest[(0<<2)+2] = src1[(0<<2)+0] * src2[(0<<2)+2] + src1[(0<<2)+1] * src2[(1<<2)+2] + src1[(0<<2)+2] * src2[(2<<2)+2] + src1[(0<<2)+3] * src2[(3<<2)+2] 
    dest[(0<<2)+3] = src1[(0<<2)+0] * src2[(0<<2)+3] + src1[(0<<2)+1] * src2[(1<<2)+3] + src1[(0<<2)+2] * src2[(2<<2)+3] + src1[(0<<2)+3] * src2[(3<<2)+3] 
    dest[(1<<2)+0] = src1[(1<<2)+0] * src2[(0<<2)+0] + src1[(1<<2)+1] * src2[(1<<2)+0] + src1[(1<<2)+2] * src2[(2<<2)+0] + src1[(1<<2)+3] * src2[(3<<2)+0] 
    dest[(1<<2)+1] = src1[(1<<2)+0] * src2[(0<<2)+1] + src1[(1<<2)+1] * src2[(1<<2)+1] + src1[(1<<2)+2] * src2[(2<<2)+1] + src1[(1<<2)+3] * src2[(3<<2)+1] 
    dest[(1<<2)+2] = src1[(1<<2)+0] * src2[(0<<2)+2] + src1[(1<<2)+1] * src2[(1<<2)+2] + src1[(1<<2)+2] * src2[(2<<2)+2] + src1[(1<<2)+3] * src2[(3<<2)+2] 
    dest[(1<<2)+3] = src1[(1<<2)+0] * src2[(0<<2)+3] + src1[(1<<2)+1] * src2[(1<<2)+3] + src1[(1<<2)+2] * src2[(2<<2)+3] + src1[(1<<2)+3] * src2[(3<<2)+3] 
    dest[(2<<2)+0] = src1[(2<<2)+0] * src2[(0<<2)+0] + src1[(2<<2)+1] * src2[(1<<2)+0] + src1[(2<<2)+2] * src2[(2<<2)+0] + src1[(2<<2)+3] * src2[(3<<2)+0] 
    dest[(2<<2)+1] = src1[(2<<2)+0] * src2[(0<<2)+1] + src1[(2<<2)+1] * src2[(1<<2)+1] + src1[(2<<2)+2] * src2[(2<<2)+1] + src1[(2<<2)+3] * src2[(3<<2)+1] 
    dest[(2<<2)+2] = src1[(2<<2)+0] * src2[(0<<2)+2] + src1[(2<<2)+1] * src2[(1<<2)+2] + src1[(2<<2)+2] * src2[(2<<2)+2] + src1[(2<<2)+3] * src2[(3<<2)+2] 
    dest[(2<<2)+3] = src1[(2<<2)+0] * src2[(0<<2)+3] + src1[(2<<2)+1] * src2[(1<<2)+3] + src1[(2<<2)+2] * src2[(2<<2)+3] + src1[(2<<2)+3] * src2[(3<<2)+3] 
    dest[(3<<2)+0] = src1[(3<<2)+0] * src2[(0<<2)+0] + src1[(3<<2)+1] * src2[(1<<2)+0] + src1[(3<<2)+2] * src2[(2<<2)+0] + src1[(3<<2)+3] * src2[(3<<2)+0] 
    dest[(3<<2)+1] = src1[(3<<2)+0] * src2[(0<<2)+1] + src1[(3<<2)+1] * src2[(1<<2)+1] + src1[(3<<2)+2] * src2[(2<<2)+1] + src1[(3<<2)+3] * src2[(3<<2)+1] 
    dest[(3<<2)+2] = src1[(3<<2)+0] * src2[(0<<2)+2] + src1[(3<<2)+1] * src2[(1<<2)+2] + src1[(3<<2)+2] * src2[(2<<2)+2] + src1[(3<<2)+3] * src2[(3<<2)+2] 
    dest[(3<<2)+3] = src1[(3<<2)+0] * src2[(0<<2)+3] + src1[(3<<2)+1] * src2[(1<<2)+3] + src1[(3<<2)+2] * src2[(2<<2)+3] + src1[(3<<2)+3] * src2[(3<<2)+3] 
    return m
}


func NewIdentityMat() Mat {
    var m Mat
    m.LoadIdentity()
    return m
}

// https://butterflyofdream.wordpress.com/2016/07/05/converting-rotation-matrices-of-left-handed-coordinate-system/
// https://en.wikipedia.org/wiki/Rotation_matrix    , right-hand, column based
// lefthand 
func NewRotXMat( rad float64 ) Mat {
    var m Mat
    m.LoadIdentity()
    m.Set( 1,1, math.Cos(rad) )
    m.Set( 2,2, math.Cos(rad) )
    m.Set(1,2,  math.Sin(rad))
    m.Set(2,1, -math.Sin(rad))
    return m
}

func NewRotYMat( rad float64 ) Mat {
    var m Mat
    m.LoadIdentity()
    m.Set( 0,0, math.Cos(rad) )
    m.Set( 2,2, math.Cos(rad))
    m.Set(0,2, -math.Sin(rad))
    m.Set(2,0,  math.Sin(rad))
    return m
}

func NewRotZMat( rad float64 ) Mat {
    var m Mat
    m.LoadIdentity()
    m.Set(0,0, math.Cos(rad))
    m.Set(1,1, math.Cos(rad))
    m.Set(0,1,  math.Sin(rad))
    m.Set(1,0, -math.Sin(rad))
    return m
}

func NewTransMat( x,y,z float64 ) Mat {
    var m Mat
    m.LoadIdentity()
    m.Set(0,3,x)
    m.Set(1,3,y)
    m.Set(2,3,z)
    return m
}

func NewProjectionMat( fovDegree, fAspectRatio, fZNear, fZFar float64 ) Mat {

    var m Mat
    m.Clear()
    /*
    // note , those code is incorrect, since fov and aspectratio inverse
    fFovRad := 1/math.Tan(  fovDegree*0.5 /180 *math.Pi  )
    m.Set(0,0, fAspectRatio * fFovRad )
    m.Set(1,1, fFovRad)
    m.Set(2,2, fZFar / (fZFar-fZNear))
    m.Set(2,3, -fZNear * fZFar / (fZFar-fZNear))
    m.Set(3,2, -1)  // negate to convert right-handed to left-handed
    m.Set(3,3, 0)
    /*/
    // for OpenGL sytle
    f := math.Abs( fZFar )
    n := math.Abs( fZNear )

    t := n * math.Tan(  fovDegree*0.5 /180 *math.Pi  )// 
    b := -t  // -
    l := -fAspectRatio * t   // -
    r := -l

    m.Set(0,0, 2*n/(r-l) ) ; m.Set(0,2, (r+l)/(r-l) )
    m.Set(1,1, 2*n/(t-b) ) ; m.Set(1,2, (t+b)/(t-b) )
    m.Set(2,2, -(f+n)/(f-n) )  ; m.Set(2,3, -(2*f*n)/(f-n) )
    m.Set(3,2, -1)
    m.Set(3,3, 0)

    // debug
    log.Println( "projection matrix test, f,n", f,n, ",l,r,t,b: ", l,r,t,b )
    for _, vec := range (   []Vec3D{ {l,b,-2,1}, {r,t,-1000,1}, {r,t,2,1}, {l,t,1000,1} } ) {
        log.Printf( "project %v -> %v" , vec , MultiplyMatrixVector( m, vec )  )
    }
    //*/
    return m
}

// GluLookAt source:  https://www.khronos.org/opengl/wiki/GluLookAt_code
func NewPointAtMat ( pos Vec3D, target Vec3D , up Vec3D ) Mat {
    newForward := target.Sub( pos ).Normalize()
    // camera, forward is mapped to -Z
    newForward.X *= -1
    newForward.Y *= -1
    newForward.Z *= -1

    a := newForward.Mul(  up.Dot( newForward )  )
    newUp := up.Sub(a).Normalize()

    newRight := newUp.Cross(  newForward ).Normalize()  // right-handed?
    // newRight := newForward.Cross(  newUp ).Normalize()

    // log.Println( newForward, newUp, newRight )

    var m Mat
    m.LoadIdentity()
    m.Set( 0,0, newRight.X )
    m.Set( 1,0, newRight.Y )
    m.Set( 2,0, newRight.Z )
    m.Set( 0,1, newUp.X )
    m.Set( 1,1, newUp.Y )
    m.Set( 2,1, newUp.Z )
    m.Set( 0,2, newForward.X )
    m.Set( 1,2, newForward.Y )
    m.Set( 2,2, newForward.Z )

    m.Set( 0,3, pos.X )
    m.Set( 1,3, pos.Y )
    m.Set( 2,3, pos.Z )

    return m
}

func QuickInverse( m0 Mat ) Mat {
    var m Mat
    m = m0  // copy

    for i:=0;i<3;i++ {
        for j:=0;j<3;j++ {
            if i != j {
                m.Set(i,j, m0.At(j,i))
            }
        }
    }
    for i:=0; i<3; i++ {
        m.Set(i,3,  -(m0.At(0,3)*m0.At(0,i) + m0.At(1,3)*m0.At(1,i) + m0.At(2,3)*m0.At(2,i))  )
    }

    return m
}

// ===============================

func (self *Mat) Clear() {
    copy( self.M[:], s16zero )
}
func (self *Mat) LoadIdentity() {
    copy( self.M[:], s16identity )
}
func (self *Mat) Set( row, col int , val float64) {
    self.M[ (row<<2) + col ] = val
}
func (self *Mat) At( row, col int ) float64 {
    return self.M[ (row<<2) + col ]
}




var s16zero []float64 = make( []float64,16 )
var s16identity []float64 = make( []float64,16 )
func init() {
    for i:=0; i<4; i++ {
        s16identity[i*4+i] = 1
    }
}

