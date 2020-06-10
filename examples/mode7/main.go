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
    "io/ioutil"
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

    sprGround = image.NewRGBA(image.Rect(0, 0, nMapSize, nMapSize))
    // debug with grid lines
    /*
    for x:=0; x<nMapSize; x+= 32 {
        for i:=-1;i<=1; i++ {
            graph.DrawLine( sprGround, x-i , 0, x-i , nMapSize , graph.COLOR_MAGENTA )
        }
    }
    for y:=0; y<nMapSize; y+= 32 {
        for i:=-1;i<=1; i++ {
            graph.DrawLine( sprGround,0,  y-i , nMapSize , y-i, graph.COLOR_BLUE )
        }
    }
    /*/
    dat, err := ioutil.ReadFile("./mariokart.spr")
    if err != nil {
        panic(err)
    }
    copy( sprGround.Pix , dat[8:] )
    //*/

    sprSky = image.NewRGBA(image.Rect(0, 0, nMapSize, nMapSize))
    dat, err = ioutil.ReadFile("./sky1.spr")
    if err != nil {
        panic(err)
    }
    copy( sprSky.Pix , dat[8:] )

    // graph.FillRect( sprGround, sprGround.Bounds(), graph.COLOR_GREEN   )
    // graph.FillRect( sprSky, sprSky.Bounds(), graph.COLOR_CYAN   )
}
func (self *MyView) Exit() {}
func (self *MyView) Update(t, dt float64) {
    window := simpleui.GetWindow()
    if simpleui.ReadKey(window, glfw.KeyA) {
        fWorldA -= 1.0 * dt
    } else if simpleui.ReadKey(window, glfw.KeyD) {
        fWorldA += 1.0 * dt
    } else if simpleui.ReadKey(window, glfw.KeyW) {
        fWorldX += math.Cos( fWorldA ) * 0.2 * dt
        fWorldY += math.Sin( fWorldA ) * 0.2 * dt
    } else if simpleui.ReadKey(window, glfw.KeyS) {
        fWorldX -= math.Cos( fWorldA ) * 0.2 * dt
        fWorldY -= math.Sin( fWorldA ) * 0.2 * dt
    } else if simpleui.ReadKey(window, glfw.KeyUp) {
        fFar += 0.1 * dt 
        if fFar > 1 {
            fFar = 1
        }
    } else if simpleui.ReadKey(window, glfw.KeyDown) {
        fFar -= 0.1 * dt 
        if fFar < 0.1 {
            fFar = 0.1
        }
    }  else if simpleui.ReadKey(window, glfw.KeyRight) {
        fNear += 0.1 * dt 
        if fNear > 1 {
            fNear = 1
        }
    } else if simpleui.ReadKey(window, glfw.KeyLeft) {
        fNear -= 0.1 * dt 
        if fNear < 0.0 {
            // fNear = 0
        }
    }

    // coordinates that represent the corners of our frustum within the map
    fFarX1 := fWorldX + math.Cos(fWorldA - fFovHalf ) * fFar
    fFarY1 := fWorldY + math.Sin(fWorldA - fFovHalf ) * fFar
    fFarX2 := fWorldX + math.Cos(fWorldA + fFovHalf ) * fFar
    fFarY2 := fWorldY + math.Sin(fWorldA + fFovHalf ) * fFar

    fNearX1 := fWorldX + math.Cos(fWorldA - fFovHalf ) * fNear
    fNearY1 := fWorldY + math.Sin(fWorldA - fFovHalf ) * fNear
    fNearX2 := fWorldX + math.Cos(fWorldA + fFovHalf ) * fNear
    fNearY2 := fWorldY + math.Sin(fWorldA + fFovHalf ) * fNear



    // draw
    graph.FillRect( self.screenImage, self.screenImage.Bounds() ,
                color.Black )

    // draw map line by line

    // when we drawing the maps, we only want them to occupy half of the screen 
    // not the full screen.
    // because this assumes that our vertical vanishing point is halfway up the screen.
    for y:=0; y<screenH/2; y++ {
        // normalize the screen height value to between 0.0 ~ 1.0
        fSampleDepth := float64(y) / (float64(screenH/2))

        if fSampleDepth == 0 {
            fSampleDepth = 0.00000001
        }
        // invert sample depth
        // scanline going across the map frustum 

        /*
        // linear
        fSampleDepth = 1 - fSampleDepth
        fStartX := (fFarX1 - fNearX1) * fSampleDepth + fNearX1
        fStartY := (fFarY1 - fNearY1) * fSampleDepth + fNearY1
        fEndX   := (fFarX2 - fNearX2) * fSampleDepth + fNearX2
        fEndY   := (fFarY2 - fNearY2) * fSampleDepth + fNearY2
        /*/
        fStartX := (fFarX1 - fNearX1) / fSampleDepth + fNearX1
        fStartY := (fFarY1 - fNearY1) / fSampleDepth + fNearY1
        fEndX   := (fFarX2 - fNearX2) / fSampleDepth + fNearX2
        fEndY   := (fFarY2 - fNearY2) / fSampleDepth + fNearY2
        //*/

        // first line is farthest away
        for x:=0; x< screenW; x++ {
            fSampleWidth := float64(x) / float64(screenW)
            // find the sample location by interpolating along that line
            fSampleX := (fEndX - fStartX)  * fSampleWidth + fStartX
            fSampleY := (fEndY - fStartY)  * fSampleWidth + fStartY

            fSampleX = sample( fSampleX )
            fSampleY = sample( fSampleY )
            if fSampleX < 0  || fSampleX >= 1 || fSampleY < 0  || fSampleY >= 1 {
                panic( "sample overflow" )
            }


            // simple sampling,  not nearest neighbor 
            sampleColor := sprGround.At(  int(math.Round( fSampleX * float64(nMapSize))) , int(math.Round( fSampleY * float64(nMapSize))) )
            // draw to bottom screen
            self.screenImage.Set( x, screenH/2 +y, sampleColor  )

            sampleColor = sprSky.At(  int(math.Round( fSampleX * float64(nMapSize))) , int(math.Round( fSampleY * float64(nMapSize))) )
            self.screenImage.Set( x, screenH/2 -y, sampleColor  )
        }
    }

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
    screenW,screenH,scale = 320,240,2
    view := NewView(screenW, screenH)
    simpleui.SetWindow( screenW, screenH, scale  )
    simpleui.Run( view )
}




var sprGround, sprSky *image.RGBA
var fWorldX, fWorldY float64 = 0.5, 0.5// Camera postion , 0~1.0
var fWorldA float64 = 0
var fNear, fFar float64  = 0.01,  0.1  // frustum
var fFovHalf = math.Pi / 4     //  the half is 45 degree  , actually 90 degree

var nMapSize  int = 1024


func sample( val float64  ) float64 {
    _ , val = math.Modf( val )
    if val < 0 {
        val += 1
    }
    if val >= 1 {
        val -= 1
    }
    return val
}
