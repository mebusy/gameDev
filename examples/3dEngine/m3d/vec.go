package m3d

import (
    "math"
)

func (self *Vec3D) Length() float64 {
    return math.Sqrt( self.Dot(*self) )
}

func (self Vec3D) Normalize() Vec3D {
    l :=  self.Length()
    return Vec3D{ self.X /l, self.Y /l, self.Z /l, 1 }
}


func (self *Vec3D) Dot( v Vec3D ) float64 {
    return self.X*v.X + self.Y*v.Y + self.Z*v.Z
}

func (self *Vec3D) Cross( v Vec3D ) Vec3D {
    var u Vec3D
    u.X = self.Y*v.Z - self.Z*v.Y
    u.Y = self.Z*v.X - self.X*v.Z
    u.Z = self.X*v.Y - self.Y*v.X
    u.W = 1
    return u
}

func (self Vec3D) NormalizeByW( ) Vec3D {
    return Vec3D{ self.X/self.W, self.Y/self.W, self.Z / self.W, 1}
}


func (self *Vec3D) Add( v Vec3D ) Vec3D {
    return Vec3D{self.X+v.X, self.Y+v.Y, self.Z+v.Z, 1 }
}

func (self *Vec3D) Sub( v Vec3D ) Vec3D {
    return Vec3D{self.X-v.X, self.Y-v.Y, self.Z-v.Z, 1 }
}
func (self *Vec3D) To( v Vec3D ) Vec3D {
    return Vec3D{v.X-self.X, v.Y-self.Y, v.Z-self.Z, 1}
}

func (self *Vec3D) Mul( k float64 ) Vec3D {
    return Vec3D{self.X*k, self.Y*k, self.Z*k, 1 }
}

func (self *Vec3D) Div( k float64 ) Vec3D {
    return Vec3D{self.X/k, self.Y/k, self.Z/k, 1 }
}







