package shadowcast

import (
    "image/draw"
    "github.com/mebusy/simpleui/graph"
)

type Edge struct {
    Sx, Sy int
    Ex, Ey int
}

type Cell struct {
    Edge_id [4]int  // -1 means not exist
    Exist bool
}

const (
    NORTH = iota
    SOUTH
    EAST
    WEST
)

type EdgePool struct {
    pool []Edge
}

func (self *EdgePool) Extend( index int, bHorizontal bool, length int )  {
    if bHorizontal {
        self.pool[index].Ex += length
    } else {
        self.pool[index].Ey += length
    }
}
func (self *EdgePool) NewEdge( sx, sy int, bHorizontal bool, length int ) int {
    var edge Edge
    edge.Sx = sx
    edge.Sy = sy
    edge.Ex = edge.Sx
    edge.Ey = edge.Sy

    if bHorizontal {
        edge.Ex += length
    } else {
        edge.Ey += length
    }

    id := len(self.pool)
    self.pool = append( self.pool, edge )
    return id
}


var edgePool EdgePool

// sx, sy: starting x,y
// w, h : along with sx,sy,which rectangle of tiles of which I wish to inspect
// pitch: is going to be set to end-world width
func ConvertTileMap2PolyMap( world []Cell,  sx,sy,w,h int, blockWidth int, pitch int ) {

    edgePool.pool = make( []Edge, 0, 64 )
    // reset old edge infos
    for x:=0; x<w; x++ {
        for y:=0; y<h; y++ {
            for j:=0; j<4; j++ {
                world[ (y+sy)*pitch + (x+sx) ].Edge_id[j] = -1
            } // end j
        } // end y
    } // end x

    // Iterate through region from top left to bottom right
    // skip the border of world , and that's simply because I don't want to have any out-of-bounds error
    for x:=1; x< w-1 ; x++ {
        for y:=1 ; y< h-1; y++ {
            // create some convenient indeices
            i := (y+sy)*pitch + (x+sx)
            n := (y+sy-1)*pitch + (x+sx)
            s := (y+sy+1)*pitch + (x+sx)
            w := (y+sy)*pitch + (x+sx-1)
            e := (y+sy)*pitch + (x+sx+1)

            // if cell exists, check if it need edges            
            if world[i].Exist {
                // if has no western neighbor, it need a western edge 
                direct := WEST
                if !world[w].Exist {
                    // it can either extend it from its northern neighbor 
                    // or it can start a new one.
                    if world[n].Edge_id[direct] != -1 {
                        // northern neighbor has a western edge , so grow it downwards
                        edgePool.Extend( world[n].Edge_id[direct] , false, blockWidth )
                        // set current cell's western edge id
                        world[i].Edge_id[direct] = world[n].Edge_id[direct]
                    }  else {
                        world[i].Edge_id[direct] = edgePool.NewEdge( (sx+x)*blockWidth, (sy+y)*blockWidth, false, blockWidth )
                    }
                }

                direct = EAST
                if !world[e].Exist {
                    // it can either extend it from its northern neighbor 
                    // or it can start a new one.
                    if world[n].Edge_id[direct] != -1 {
                        // northern neighbor has a western edge , so grow it downwards
                        edgePool.Extend( world[n].Edge_id[direct] , false, blockWidth )
                        // set current cell's western edge id
                        world[i].Edge_id[direct] = world[n].Edge_id[direct]
                    } else {
                        // only difference from western edge checking , is that edge.Sx should + 1*blockWidth
                        world[i].Edge_id[direct] = edgePool.NewEdge( (sx+x+1)*blockWidth, (sy+y)*blockWidth, false, blockWidth )
                    }
                }

                direct = NORTH
                if !world[n].Exist {
                    if world[w].Edge_id[direct] != -1 {
                        edgePool.Extend( world[w].Edge_id[direct] , true, blockWidth )
                        world[i].Edge_id[direct] = world[w].Edge_id[direct]
                    }  else {
                        world[i].Edge_id[direct] = edgePool.NewEdge( (sx+x)*blockWidth, (sy+y)*blockWidth, true, blockWidth )
                    }
                }

                direct = SOUTH
                if !world[s].Exist {
                    if world[w].Edge_id[direct] != -1 {
                        edgePool.Extend( world[w].Edge_id[direct] , true, blockWidth )
                        world[i].Edge_id[direct] = world[w].Edge_id[direct]
                    }  else {
                        // only difference from northern edge checking , is that edge.Sy should + 1*blockWidth
                        world[i].Edge_id[direct] = edgePool.NewEdge( (sx+x)*blockWidth, (sy+y+1)*blockWidth, true, blockWidth )
                    }
                }

            }
        } // end y
    } // end x
}

func DrawPolyonMap( dst draw.Image ) {
    if edgePool.pool == nil {
        return
    }
    for _, e := range edgePool.pool {
        graph.DrawLine( dst, e.Sx , e.Sy, e.Ex, e.Ey ,  graph.COLOR_RED  )
    }
}



