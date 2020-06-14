package m3d

import (
)

func MultiplyMatrixVector( m Mat, i Vec3D,  o *Vec3D ) {
    o.X = i.X*m.M[(0<<2)+0]+i.Y*m.M[(0<<2)+1]+i.Z*m.M[(0<<2)+2]+m.M[(0<<2)+3]
    o.Y = i.X*m.M[(1<<2)+0]+i.Y*m.M[(1<<2)+1]+i.Z*m.M[(1<<2)+2]+m.M[(1<<2)+3]
    o.Z = i.X*m.M[(2<<2)+0]+i.Y*m.M[(2<<2)+1]+i.Z*m.M[(2<<2)+2]+m.M[(2<<2)+3]
    w  := i.X*m.M[(3<<2)+0]+i.Y*m.M[(3<<2)+1]+i.Z*m.M[(3<<2)+2]+m.M[(3<<2)+3]
    // if w == 0, this point makes no sense
    // this is internally done by openGL
    if w != 0.0 {
        o.X /= w
        o.Y /= w
        o.Z /= w
    }
}



