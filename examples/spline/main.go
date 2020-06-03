package main

import (
    "github.com/mebusy/simpleui"
    "github.com/go-gl/glfw/v3.1/glfw"
    "image"
    "image/color"
    // "image/draw"
    "github.com/mebusy/simpleui/graph"
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
    // clear
    graph.FillRect( self.screenImage , self.screenImage.Bounds(), color.Black )

    spline_path.Update(dt)
    spline_path.Draw( self.screenImage )
}

func (self *MyView) SetAudioDevice(audio *simpleui.Audio) {}
func (self *MyView) OnKey(key glfw.Key) {
    switch key {
    case glfw.KeyX:
        switchControlPoint()
    }
}
func (self *MyView) TextureBuff() []uint8 {
    return self.screenImage.Pix
}
func (self *MyView) Title() string {
    return "my game"
}


func main() {
    w,h,scale := 128,80,8
    view := NewView(w,h)
    simpleui.SetWindow( w,h, scale  )
    simpleui.Run( view )
}


