package main

func (self *Spline) getSplinePoint( t float64, bLoop bool ) Point2D {
    // index of ctl_points
    var p0,p1,p2,p3 int
    if !bLoop {
        p1 = int(t) + 1
        p2 = p1 + 1
        p3 = p2 + 1
        p0 = p1 - 1
    } else {
        p1 = int(t)
        p2 = (p1 + 1) % len( self.ctl_points )
        p3 = (p2 + 1) % len( self.ctl_points )
        if p1 >= 1 {
            p0 = p1 -1
        } else {
            p0 = len( self.ctl_points )-1
        }
    }
    pt_indices := []int {
        p0,p1,p2,p3,
    }

    // when the number of control ctl_points > 4
    // t will be > 1
    t = t - float64(int(t))

    tt := t*t
    ttt := tt*t

    q0 := -ttt + 2*tt - t
    q1 := 3*ttt - 5*tt + 2
    q2 := -3*ttt + 4*tt + t
    q3 := ttt - tt
    pt_values := []float64 {
        q0,q1,q2,q3,
    }

    var tx, ty float64
    for i:=0; i<len(pt_indices); i++ {
        tx += self.ctl_points[ pt_indices[i] ].x * pt_values[i]
        ty += self.ctl_points[ pt_indices[i] ].y * pt_values[i]
    }
    tx *= 0.5
    ty *= 0.5
    return Point2D{tx,ty}
}

func (self *Spline) getSplineSlope( t float64, bLoop bool ) Point2D {
    // index of ctl_points
    var p0,p1,p2,p3 int
    if !bLoop {
        p1 = int(t) + 1
        p2 = p1 + 1
        p3 = p2 + 1
        p0 = p1 - 1
    } else {
        p1 = int(t)
        p2 = (p1 + 1) % len( self.ctl_points )
        p3 = (p2 + 1) % len( self.ctl_points )
        if p1 >= 1 {
            p0 = p1 -1
        } else {
            p0 = len( self.ctl_points )-1
        }
    }
    pt_indices := []int {
        p0,p1,p2,p3,
    }

    // when the number of control ctl_points > 4
    // t will be > 1
    t = t - float64(int(t))

    tt := t*t

    q0 := -3*tt + 2*2*t - 1
    q1 := 3*3*tt - 5*2*t
    q2 := -3*3*tt + 4*2*t + 1
    q3 := 3*tt - 2*t
    pt_values := []float64 {
        q0,q1,q2,q3,
    }

    var tx, ty float64
    for i:=0; i<len(pt_indices); i++ {
        tx += self.ctl_points[ pt_indices[i] ].x * pt_values[i]
        ty += self.ctl_points[ pt_indices[i] ].y * pt_values[i]
    }
    tx *= 0.5
    ty *= 0.5
    return Point2D{tx,ty}
}

// this is a naive solution
func (self *Spline) CalculateSegmentLength(node int , bLoop bool) float64 {
    var fLength float64 = 0.0
    var fStepSize float64 = 0.005

    // var old_pt, new_pt Point2D
    old_pt := self.getSplinePoint( float64(node), bLoop )

    var t float64
    for t=0; t<1; t+= fStepSize {
        new_pt := self.getSplinePoint( float64(node)+t, bLoop )
        fLength += old_pt.DistanceTo( new_pt)
        old_pt = new_pt
    }

    return fLength
}

// convert a real path distance to a normalized distance
//  e.g.  total length 136,  GetNormalizedOffset(90) => 6.8
// postion p: from the real world
//     it will be some positon along the real length of the spline
func (self *Spline) GetNormalizedOffset( p float64 ) float64 {
    // work out which node is the base.
    i := 0
    for p > self.ctl_pt_lengths[i] {
        p -= self.ctl_pt_lengths[i]
        i++
    }
    // the fractional is the offset
    return float64(i) + p/self.ctl_pt_lengths[i]
}

