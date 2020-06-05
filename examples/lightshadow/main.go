package main

import (
	"image"
	"image/color"

	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/mebusy/simpleui"

	"github.com/mebusy/simpleui/graph"
	// "image/draw"
    "shadowcast"
    "log"
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
}
func (self *MyView) Exit() {}
func (self *MyView) Update(t, dt float64) {
    window := simpleui.GetWindow()

    mx,my := simpleui.GetCursorPosInWindow(window)
    if simpleui.ReadMouse( window , glfw.MouseButtonLeft ) {
        idx_cell := int(my)/CELL_WIDTH * nWorldWidth  +  int(mx)/CELL_WIDTH
        log.Println( idx_cell )
    }

    graph.FillRect( self.screenImage, self.screenImage.Bounds() ,
                color.Black )

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



