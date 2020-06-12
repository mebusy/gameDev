
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



