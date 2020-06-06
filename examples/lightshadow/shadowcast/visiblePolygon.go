package shadowcast

import (
    "fmt"
    "image/color"
    "image/draw"
    "math"
    "sort"
    "github.com/mebusy/simpleui/graph"
    // "log"
)

type VisibleVertex struct {
    theta float64
    x, y int
}

type VisiblePolygonPoints struct {
    points  []VisibleVertex
}

func (self *VisiblePolygonPoints) Add( theta float64,  x,y int ) {
    self.points = append( self.points , VisibleVertex{theta,x,y } )
}
func (self *VisiblePolygonPoints) Reorder( ) {
    sort.SliceStable(self.points, func(i, j int) bool {
        return self.points[i].theta < self.points[j].theta
    })
}

var visiblePolygonPoints VisiblePolygonPoints

func CalculateVisibilityPolygon( ox,oy int, radius float64 ) {
    visiblePolygonPoints.points = make( []VisibleVertex,0,64 )
    m_dup := make(map[string]bool)
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
                    key :=  fmt.Sprintf( "%d,%d", int(min_px), int(min_py) ) 
                    if !m_dup[key] {
                        visiblePolygonPoints.Add( min_theta, int(min_px), int(min_py) )
                        m_dup[key] = true
                    }
                }
            } // end for j < 3
        } // end for i < 2
    } // end for e1
    // Sort perimeter points by angle from source. This will allow
    // us to draw a triangle fan.
    visiblePolygonPoints.Reorder()
    // log.Println( "nRay:" ,  len(visiblePolygonPoints.points) )
}

func DrawPolygonVisible( dst draw.Image , sx, sy int,  color color.Color ) {
    triangle := graph.NewTriangle(0,0,0,0,0,0)
    points := visiblePolygonPoints.points
    nPoint := len(points)
    // log.Println( nPoint,  points )
    for i:=0; i<nPoint; i++ {
        triangle.SetVert(0, sx,sy)
        triangle.SetVert(1, points[i].x, points[i].y)
        idx := (i+1) % nPoint
        triangle.SetVert(2, points[idx].x, points[idx].y)
        graph.FillTriangle( dst, triangle, color )
    }
}
