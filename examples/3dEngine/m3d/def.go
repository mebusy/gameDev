package m3d

// import "math"

type Vec3D struct {
    X,Y,Z  float64
    W float64
}

type Triangle struct {
    P [3]Vec3D
    Color float64
}

type Mesh struct {
    Tris []Triangle
}

type Mat struct {
    M [16]float64
}

// ===============================


// ===============================

func (self *Triangle) CalculateNormal() Vec3D {
    var normal, line1, line2 Vec3D
    // P[0], P[1], p[2] in clockwise
    line1.X = self.P[1].X - self.P[0].X
    line1.Y = self.P[1].Y - self.P[0].Y
    line1.Z = self.P[1].Z - self.P[0].Z
    line2.X = self.P[2].X - self.P[0].X
    line2.Y = self.P[2].Y - self.P[0].Y
    line2.Z = self.P[2].Z - self.P[0].Z
    normal = line1.Cross(line2)
    return normal.Normalize()
}

func (self *Triangle) MidPointZ() float64 {
    return (self.P[0].Z + self.P[1].Z + self.P[2].Z)/3
}

