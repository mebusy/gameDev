package main

import (
	"image"
	"image/color"

	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/mebusy/simpleui"

	"github.com/mebusy/simpleui/graph"
	// "image/draw"
)

type MyView struct {
    screenImage *image.RGBA
}

func NewView( w,h int) *MyView {
    view := &MyView{}
    view.screenImage = image.NewRGBA(image.Rect(0, 0, w, h))
    return view
}

func (self *MyView) Enter() {}
func (self *MyView) Exit() {}
func (self *MyView) Update(t, dt float64) {

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
