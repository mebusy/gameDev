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
    "math"
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

    // we even don't need make our track sense
    vecTrack = make([]Track,0, 16)
    // short section for start/finish line
    vecTrack = append( vecTrack , Track{ 0,10 }  )

    vecTrack = append( vecTrack , Track{ 0,200 }  )
    // full curvature 1, to the right.
    vecTrack = append( vecTrack , Track{ 1,200 }  )
    vecTrack = append( vecTrack , Track{ 0,400 }  )
    vecTrack = append( vecTrack , Track{ -1,100 }  )
    vecTrack = append( vecTrack , Track{ 0,200 }  )
    vecTrack = append( vecTrack , Track{ -1,200 }  )
    vecTrack = append( vecTrack , Track{ 1,200 }  )
    vecTrack = append( vecTrack , Track{ 0.2,500 }  )
    vecTrack = append( vecTrack , Track{ 0,200 }  )
}
func (self *MyView) Exit() {}
func (self *MyView) Update(t, dt float64) {
    window := simpleui.GetWindow()
    if simpleui.ReadKey( window, glfw.KeyUp ) {
        fCarDistance += 100 * dt
    } else if simpleui.ReadKey( window, glfw.KeyDown ) {
        fCarDistance -= 100 * dt
    }

    graph.FillRect( self.screenImage, self.screenImage.Bounds() ,
                color.Black )

    // halfScreenW := screenW/2
    halfScreenH := screenH/2

    // draw track
    for y:=0; y<halfScreenH; y++ {
        // depth
        fPerspective := float64(y)/float64(halfScreenH)
        // we convert screen into a normalized space [0,1]
        fMiddlePoint := 0.5
        // the road occupies 60% of the screen
        fRoadWidth := fPerspective * 0.8 + 0.1
        // clipboard, the red/which line going around the track
        fClipWidth := fRoadWidth * 0.15
        // road track is symmetric 
        fRoadWidth *= 0.5

        nRow := halfScreenH + y

        bDarkGrass := math.Sin(20* math.Pow(1-fPerspective, 3) +fCarDistance*0.1) < 0
        bRedClip :=   math.Sin(80* math.Pow(1-fPerspective, 2) +fCarDistance    ) < 0

        for x:=0; x<screenW; x++ {

            nLeftGrass := int((fMiddlePoint - fRoadWidth - fClipWidth)*float64(screenW))
            nLeftClip := int((fMiddlePoint - fRoadWidth )*float64(screenW))
            nRightGrass := int((fMiddlePoint + fRoadWidth + fClipWidth)*float64(screenW))
            nRightClip := int((fMiddlePoint + fRoadWidth )*float64(screenW))


            grassColor := graph.COLOR_GREEN
            if bDarkGrass {
                grassColor = graph.COLOR_DARKGREEN
            }
            clipColor := graph.COLOR_WHITE
            if bRedClip {
                clipColor = graph.COLOR_RED
            }

            if x < nLeftGrass || x > nRightGrass{
                self.screenImage.Set(x, nRow, grassColor)
            } else  if (x >= nLeftGrass && x < nLeftClip) || (x > nRightClip && x <= nRightGrass ) {
                self.screenImage.Set(x, nRow, clipColor)
            } else {
                self.screenImage.Set(x, nRow, graph.COLOR_GRAY)
            }

        }
    }


    // draw car
    // carpostion:   -1,0,+1 ,  is going to be 0 when it is at the middle of track
    // nCarPos := halfScreenW + int(float64(halfScreenW) * fCarPos)

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
var fCarDistance float64 = 0.0

type Track struct {
    Curvature float64
    Distance float64
}
var vecTrack []Track

