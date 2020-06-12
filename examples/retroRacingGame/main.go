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

    for _ , track := range vecTrack {
        fTotalTrackDistance += track.Distance
    }
}
func (self *MyView) Exit() {}
func (self *MyView) Update(t, dt float64) {
    window := simpleui.GetWindow()
    if simpleui.ReadKey( window, glfw.KeyUp ) {
        fCarSpeed += 2.0*dt
    } else {
        fCarSpeed -= 1.0*dt
    }

    if simpleui.ReadKey( window, glfw.KeyLeft ) {
        fPlayerCurvature -= 0.7 * dt
    } else if simpleui.ReadKey( window, glfw.KeyRight ) { 
        fPlayerCurvature += 0.7 * dt
    }

    if math.Abs(fPlayerCurvature - fAccTrackCurvature)  >= 0.8 {
        // car gone off the track, slow down
        fCarSpeed -= 5 * dt
    }
    // final clamp speed
    if fCarSpeed < 0 {
        fCarSpeed = 0
    } else if fCarSpeed > 1 {
        fCarSpeed = 1
    }

    // move car along track according to car speed
    fCarDistance += (70*fCarSpeed) * dt
    if fCarDistance >= fTotalTrackDistance {
        fCarDistance -= fTotalTrackDistance
    }

    // Get Point on Track
    tmpCarDistance := 0.0
    nTrackSection := 0
    // find postion on track ( naive , not optimal )
    for nTrackSection < len( vecTrack ) && tmpCarDistance <= fCarDistance {
        tmpCarDistance += vecTrack[nTrackSection].Distance
        nTrackSection ++
    }

    // change curvature to target curvature little by little
    fTargetCurvature := vecTrack[nTrackSection - 1].Curvature
    fTrackCurveDiff := (fTargetCurvature - fCurvature) * dt * fCarSpeed
    fCurvature += fTrackCurveDiff

    fAccTrackCurvature +=  (fCurvature) * dt * fCarSpeed

    graph.FillRect( self.screenImage, self.screenImage.Bounds() ,
                color.Black )

    halfScreenW := screenW/2
    halfScreenH := screenH/2


    // draw sky
    for y:=0; y< halfScreenH; y++ {
        if y < halfScreenH/2 {
            graph.DrawLine( self.screenImage,0,y,screenW,y, graph.COLOR_BLUE )
        } else {
            graph.DrawLine( self.screenImage,0,y,screenW,y, graph.COLOR_DARKBLUE )
        }
    }

    // draw Scenery -- out hills are a rectified sine wave, where the phase 
    // accumulated track curvature
    for x:=0; x<screenW; x++ {
        nHillHeight := int(math.Abs( math.Sin( float64(x)*0.01 + fAccTrackCurvature)*16 ))
        // println( x, nHillHeight  )
        graph.DrawLine(self.screenImage, x, halfScreenH - nHillHeight ,x , halfScreenH , graph.COLOR_DARKYELLOW  )
    }

    bCheckPoint := nTrackSection - 1 == 0
    roadColor := graph.COLOR_GRAY
    if bCheckPoint {
        roadColor = graph.COLOR_WHITE
    }

    // draw track
    for y:=0; y<halfScreenH; y++ {
        // depth
        fPerspective := float64(y)/float64(halfScreenH)
        // we convert screen into a normalized space [0,1]
        fMiddlePoint := 0.5 + fCurvature * math.Pow( 1-fPerspective, 3 )
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
                self.screenImage.Set(x, nRow, roadColor )
            }

        }
    }


    // draw car
    // carpostion:   -1,0,+1 ,  is going to be 0 when it is at the middle of track
    fCarPos := fPlayerCurvature - fAccTrackCurvature
    nCarPos := halfScreenW + int(float64(halfScreenW) * fCarPos)
    carW := 20
    carH := 12
    graph.DrawRectangle( self.screenImage, nCarPos- carW/2 , 80 , nCarPos + carW/2 , 80 + carH, graph.COLOR_BLUE )
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

var fCarSpeed float64 = 0.0
var fCarDistance float64 = 0.0
var fCurvature float64 = 0.0  // for curvature displaying

var fAccTrackCurvature float64  // acc horizontal offset
var fPlayerCurvature float64

var fTotalTrackDistance float64

type Track struct {
    Curvature float64
    Distance float64
}
var vecTrack []Track

