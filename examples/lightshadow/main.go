package main

import (
	"image"
	"image/color"

	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/mebusy/simpleui"

	"github.com/mebusy/simpleui/graph"
	// "image/draw"
    "shadowcast"
    // "log"
)

type MyView struct {
    screenImage *image.RGBA
}

func NewView( w,h int) *MyView {
    view := &MyView{}
    view.screenImage = image.NewRGBA(image.Rect(0, 0, w, h))
    return view
}

func (self *MyView) Enter() {
    window := simpleui.GetWindow()
    window.SetMouseButtonCallback( func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {

            if action == glfw.Press && button == glfw.MouseButtonLeft {
                mx,my := simpleui.GetCursorPosInWindow(window)
                idx_cell := int(my)/CELL_WIDTH * nWorldWidth  +  int(mx)/CELL_WIDTH
                world[idx_cell].Exist = !world[idx_cell].Exist
                // re-generate poly map
                shadowcast.ConvertTileMap2PolyMap( world[:] , 0,0,nWorldWidth, nWorldHeight, CELL_WIDTH, nWorldWidth  )
            }
        },
    )
}
func (self *MyView) Exit() {}
func (self *MyView) Update(t, dt float64) {
    // window := simpleui.GetWindow()

    graph.FillRect( self.screenImage, self.screenImage.Bounds() ,
                color.Black )

    for y :=0; y<nWorldHeight; y++ {
        for x :=0; x<nWorldWidth ; x++ {
            idx_cell := y * nWorldWidth  + x
            if world[idx_cell].Exist {
                bound := image.Rect( x*CELL_WIDTH, y*CELL_WIDTH, x*CELL_WIDTH+CELL_WIDTH, y*CELL_WIDTH+CELL_WIDTH )
                graph.FillRect( self.screenImage, bound ,  graph.COLOR_BLUE )
            }
        }
    }


    shadowcast.DrawEdge( self.screenImage )

}

func (self *MyView) OnKey(key glfw.Key) {}
func (self *MyView) TextureBuff() []uint8 {
    return self.screenImage.Pix
}
func (self *MyView) Title() string {
    return "my game"
}


func main() {
    w,h,scale := 640,480,1
    view := NewView(w,h)
    simpleui.SetWindow( w,h, scale  )
    simpleui.Run( view )
}

// ===============================================

const nWorldWidth = 40
const nWorldHeight = 30
const CELL_WIDTH = 16
var world [nWorldWidth * nWorldHeight]shadowcast.Cell



