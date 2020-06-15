package m3d

import "math"

type Vec3D struct {
    X,Y,Z  float64
}

type Triangle struct {
    P [3]Vec3D
}

type Mesh struct {
    Tris []Triangle
}

type Mat struct {
    M [16]float64
}

// ===============================

func (self *Vec3D) Normalize() {
    l := math.Sqrt(  self.X*self.X + self.Y*self.Y + self.Z*self.Z )
    self.X /= l
    self.Y /= l
    self.Z /= l
}

func (self *Vec3D) VectorTo( to Vec3D  ) Vec3D {
    return Vec3D{ to.X-self.X, to.Y-self.Y, to.Z-self.Z }
}

func (self *Vec3D) Dot( v Vec3D ) float64 {
    return self.X*v.X + self.Y*v.Y + self.Z*v.Z
}

// ===============================

func (self *Triangle) CalculateNormal() Vec3D {
    var normal, line1, line2 Vec3D
    line1.X = self.P[1].X - self.P[0].X
    line1.Y = self.P[1].Y - self.P[0].Y
    line1.Z = self.P[1].Z - self.P[0].Z
    line2.X = self.P[2].X - self.P[0].X
    line2.Y = self.P[2].Y - self.P[0].Y
    line2.Z = self.P[2].Z - self.P[0].Z

    normal.X = line1.Y*line2.Z - line1.Z*line2.Y
    normal.Y = line1.Z*line2.X - line1.X*line2.Z
    normal.Z = line1.X*line2.Y - line1.Y*line2.X
    normal.Normalize()
    return normal
}

// ===============================

func (self *Mat) Clear() {
    copy( self.M[:], s16zero )
}
func (self *Mat) Identity() {
    copy( self.M[:], s16identity )
}
func (self *Mat) Set( row, col int , val float64) {
    self.M[ (row<<2) + col ] = val
}
/*
func (self *Mat) At( row, col int ) float64 {
    return self.M[ (row<<2) + col ]
}
//*/




var s16zero []float64 = make( []float64,16 )
var s16identity []float64 = make( []float64,16 )
func init() {
    for i:=0; i<4; i++ {
        s16identity[i*4+i] = 1
    }
}

