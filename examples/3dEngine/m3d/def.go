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

