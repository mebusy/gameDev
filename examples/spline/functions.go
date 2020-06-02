package main

func (self *Spline) getSplinePoint( t float32, bLoop bool ) Point2D {
    // index of points
    var p0,p1,p2,p3 int
    if !bLoop {
        p1 = int(t) + 1
        p2 = p1 + 1
        p3 = p2 + 1
        p0 = p1 - 1
    } else {
        p1 = int(t)
        p2 = (p1 + 1) % len( self.points )
        p3 = (p2 + 1) % len( self.points )
        if p1 >= 1 {
            p0 = p1 -1
        } else {
            p0 = len( self.points )-1
        }
    }
    pt_indices := []int {
        p0,p1,p2,p3,
    }

    // when the number of control points > 4
    // t will be > 1
    t = t - float32(int(t))

    tt := t*t
    ttt := tt*t

    q0 := -ttt + 2*tt - t
    q1 := 3*ttt - 5*tt + 2
    q2 := -3*ttt + 4*tt + t
    q3 := ttt - tt
    pt_values := []float32 {
        q0,q1,q2,q3,
    }

    var tx, ty float32
    for i:=0; i<len(pt_indices); i++ {
        tx += self.points[ pt_indices[i] ].x * pt_values[i]
        ty += self.points[ pt_indices[i] ].y * pt_values[i]
    }
    tx *= 0.5
    ty *= 0.5
    return Point2D{tx,ty}
}

func (self *Spline) getSplineSlope( t float32, bLoop bool ) Point2D {
    // index of points
    var p0,p1,p2,p3 int
    if !bLoop {
        p1 = int(t) + 1
        p2 = p1 + 1
        p3 = p2 + 1
        p0 = p1 - 1
    } else {
        p1 = int(t)
        p2 = (p1 + 1) % len( self.points )
        p3 = (p2 + 1) % len( self.points )
        if p1 >= 1 {
            p0 = p1 -1
        } else {
            p0 = len( self.points )-1
        }
    }
    pt_indices := []int {
        p0,p1,p2,p3,
    }

    // when the number of control points > 4
    // t will be > 1
    t = t - float32(int(t))

    tt := t*t

    q0 := -3*tt + 2*2*t - 1
    q1 := 3*3*tt - 5*2*t
    q2 := -3*3*tt + 4*2*t + 1
    q3 := 3*tt - 2*t
    pt_values := []float32 {
        q0,q1,q2,q3,
    }

    var tx, ty float32
    for i:=0; i<len(pt_indices); i++ {
        tx += self.points[ pt_indices[i] ].x * pt_values[i]
        ty += self.points[ pt_indices[i] ].y * pt_values[i]
    }
    tx *= 0.5
    ty *= 0.5
    return Point2D{tx,ty}
}

