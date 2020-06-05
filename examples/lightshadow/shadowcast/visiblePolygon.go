package shadowcast

import (
    "math"
)

type VisibleVertex struct {
    theta float64
    x, y float64
}

type VisiblePolygonPoints struct {
    points  []VisibleVertex
}

func (self *VisiblePolygonPoints) Add( theta float64,  x,y float64 ) {
    self.points = append( self.points , VisibleVertex{theta,x,y } )
}

var visiblePolygonPoints VisiblePolygonPoints

func CalculateVisibilityPolygon( ox,oy int, radius float64 ) {
    visiblePolygonPoints.points = make( []VisibleVertex,0,64 )

    for _, e1 := range edgePool.pool {
        // take the start point, then the end point
        coords := []int{ e1.Sx,e1.Sy,e1.Ex,e1.Ey }
        for i:=0; i<2; i++ {
            rdx := coords[i*2] - ox
            rdy := coords[i*2+1] - oy
            base_ang := math.Atan2( float64(rdy), float64(rdx) )

            for j:=0; j<3;j++ {
                theta := base_ang - float64(1-j)*0.0001
                // create ray along angle for required distance
                rdx := radius * math.Cos( theta )
                rdy := radius * math.Sin( theta )

                var min_t1 float64 = math.Inf(1)
                var min_px, min_py float64
                var min_theta float64
                var bValid bool

                // check for ray intersection with all edges
                for _, e2 := range edgePool.pool {
                    // create line segment vector
                    sdx := float64( e2.Ex - e2.Sx )
                    sdy := float64( e2.Ey - e2.Sy )
                    // make sure they are efficiently different
                    // for example both have dy == 0 , 
                    // then both ray and edge could be horizontal.
                    if math.Abs( sdx-rdx )>0 && math.Abs( sdy-rdy )>0 {
                        // t2 is normalised distance from line segment start to line segment end of intersect point
                        t2 := (rdx * float64(e2.Sy - oy) + (rdy * float64(ox - e2.Sx))) / (sdx * rdy - sdy * rdx)
                        // t1 is normalised distance from source along ray to ray length of intersect point
                        t1 := (float64(e2.Sx) + sdx * t2 - float64(ox)) / rdx
                        // If intersect point exists along ray, and along line
                        // segment then intersect point is valid
                        if t1 > 0 && t2 >= 0 && t2 <= 1.0 {
                            if t1 < min_t1 {
                                min_t1 = t1
                                min_px = float64(ox) + rdx * t1
                                min_py = float64(oy) + rdy * t1
                                min_theta = math.Atan2(min_py - float64(oy), min_px - float64(ox))
                                bValid = true
                            }
                        }
                    } // end if Abs > 0
                } // end for e2

                if bValid {
                    visiblePolygonPoints.Add( theta, min_px, min_py )
                }
            } // end for j < 3
        } // end for i < 2
    } // end for e1
    // Sort perimeter points by angle from source. This will allow
    // us to draw a triangle fan.
    
}



// https://stackoverflow.com/questions/563198/how-do-you-detect-where-two-line-segments-intersect
// Returns 1 if the lines intersect, otherwise 0. In addition, if the lines
// intersect the intersection point may be stored in the floats i_x and i_y.
func get_line_intersection( p0_x,  p0_y,  p1_x,  p1_y,
     p2_x,  p2_y,  p3_x,  p3_y float64) (bool,float64,float64) {

    var i_x, i_y float64

    var s1_x, s1_y, s2_x, s2_y float64
    s1_x = p1_x - p0_x;     s1_y = p1_y - p0_y;
    s2_x = p3_x - p2_x;     s2_y = p3_y - p2_y;

    var s, t float64
    s = (-s1_y * (p0_x - p2_x) + s1_x * (p0_y - p2_y)) / (-s2_x * s1_y + s1_x * s2_y)
    t = ( s2_x * (p0_y - p2_y) - s2_y * (p0_x - p2_x)) / (-s2_x * s1_y + s1_x * s2_y)

    if s >= 0 && s <= 1 && t >= 0 && t <= 1 {
        // Collision detected
        i_x = p0_x + (t * s1_x)
        i_y = p0_y + (t * s1_y)
        return true, i_x, i_y
    }

    return false, 0.0, 0.0
}
