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

func (self *MyView) Enter() {
    rand.Seed( time.Now().Unix() )

    nOutputSize = screenW
    fNoiseSeed1D = make([]float64, 256)
    fPerlinNoise1D = make([]float64, 256)

    GenerateSeedNoise()
    GeneratePerlinNoise1D(  nOutputSize , fNoiseSeed1D, fPerlinNoise1D  )

}
func (self *MyView) Exit() {}
func (self *MyView) Update(t, dt float64) {

    graph.FillRect( self.screenImage, self.screenImage.Bounds() ,
                color.Black )

    for x:=0; x<nOutputSize; x++ {
        // minus y effectively make y grows from bottom to top
        y := -int(fPerlinNoise1D[x] * float64(screenH/2)) + (screenH)/2
        graph.DrawLine( self.screenImage , x, y , x , screenH/2 , graph.COLOR_GREEN )
    }
}

func (self *MyView) OnKey(key glfw.Key) {
    if key == glfw.KeyR {
        GenerateSeedNoise()
        GeneratePerlinNoise1D(  nOutputSize , fNoiseSeed1D, fPerlinNoise1D  )
    }
}
func (self *MyView) TextureBuff() []uint8 {
    return self.screenImage.Pix
}
func (self *MyView) Title() string {
    return "my game"
}


var screenW, screenH, scale int

func main() {
    screenW,screenH,scale = 256,256,2
    view := NewView(screenW,screenH)
    simpleui.SetWindow( screenW,screenH , scale  )
    simpleui.Run( view )
}

var fNoiseSeed1D []float64
var fPerlinNoise1D []float64
var nOutputSize int

func GeneratePerlinNoise1D( nCount  int, fSeed []float64, fOutput []float64 ) {
    fOctaves := math.Log2(  float64(nCount ) )
    nOctaves := int(fOctaves)
    if fOctaves != float64(nOctaves) {
        panic( "output count must be power of 2" )
    }
    fScaleAcc := 2 - 1/float64(nCount/2)
    for x:=0; x<nCount ; x++ {
        fNoise := 0.0
        fScale := 1.0
        for o:=0; o< nOctaves; o++ {
            nPitch := nCount  >> o
            nSample1 := x &^(nPitch-1)  //  clamp left (left/right sometimes inverse)
            nSample2 := (nSample1 + nPitch) % nCount  // clamp right 
            // how far into the pitch
            fBlend   := float64(x - nSample1) / float64(nPitch)
            // interpolation    nSample1:0<---fBlend---> 1:nSample2
            fSample := (1-fBlend)*fSeed[nSample1] + fBlend * fSeed[nSample2]
            // fNoise is going to accumulate the noise for location X within our output array. 
            fNoise += fSample * fScale
            // halve scale for next octave
            fScale *= 0.5
        }
        // output , make sure it lies between [0,1]
        fOutput[x] = fNoise / fScaleAcc
    }
}

func GenerateSeedNoise() {
    for i:=0; i<nOutputSize; i++ {
        fNoiseSeed1D[i] = rand.Float64()
    }
}
