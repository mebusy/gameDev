
package m3d

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

