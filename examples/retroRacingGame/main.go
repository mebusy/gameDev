package main

import (
    "github.com/mebusy/simpleui"
    "github.com/go-gl/glfw/v3.1/glfw"
    "image"
    "image/color"
    "github.com/mebusy/simpleui/graph"
    // "image/draw"
    "math/rand"
    // "log"
    "time"
)

type MyView struct {
    screenImage *image.RGBA
}

func NewView( w,h int) *MyView {
    view := &MyView{}
    view.screenImage = image.NewRGBA(image.Rect(0, 0, w, h))
    return view
}

var test_pts [6]int
func (self *MyView) Enter() {
    rand.Seed( time.Now().Unix() )
}
func (self *MyView) Exit() {}
func (self *MyView) Update(t, dt float64) {

    graph.FillRect( self.screenImage, self.screenImage.Bounds() ,
                color.Black )

    halfScreenW := screenW/2
    halfScreenH := screenH/2

    // draw track
    for y:=0; y<halfScreenH; y++ {
        for x:=0; x<screenW; x++ {
            // we convert screen into a normalized space [0,1]
            fMiddlePoint := 0.5
            // the road occupies 60% of the screen
            fRoadWidth := 0.6
            // clipboard, the red/which line going around the track
            fClipWidth := fRoadWidth * 0.15
            // road track is symmetric 
            fRoadWidth *= 0.5

            nLeftGrass := int((fMiddlePoint - fRoadWidth - fClipWidth)*float64(screenW))
            nLeftClip := int((fMiddlePoint - fRoadWidth )*float64(screenW))
            nRightGrass := int((fMiddlePoint + fRoadWidth + fClipWidth)*float64(screenW))
            nRightClip := int((fMiddlePoint + fRoadWidth )*float64(screenW))

            nRow := halfScreenH + y
            if x < nLeftGrass || x > nRightGrass{
                self.screenImage.Set(x, nRow, graph.COLOR_GREEN)
            } else  if (x >= nLeftGrass && x < nLeftClip) || (x > nRightClip && x <= nRightGrass ) {
                self.screenImage.Set(x, nRow, graph.COLOR_RED)
            } else {
                self.screenImage.Set(x, nRow, graph.COLOR_GRAY)
            }

        }
    }


    // draw car
    // carpostion:   -1,0,+1 ,  is going to be 0 when it is at the middle of track
    nCarPos := halfScreenW + int(float64(halfScreenW) * fCarPos)

}

func (self *MyView) OnKey(key glfw.Key) {}
func (self *MyView) TextureBuff() []uint8 {
    return self.screenImage.Pix
}
func (self *MyView) Title() string {
    return "my game"
}


var screenW, screenH, scale int

func main() {
    screenW,screenH,scale = 160,100,4
    view := NewView(screenW,screenH)
    simpleui.SetWindow( screenW,screenH , scale  )
    simpleui.Run( view )
}

var fCarPos float64 = 0.0
