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

    nOutputSize := screenW
    fNoiseSeed1D = make([]float64, nOutputSize)
    fPerlinNoise1D = make([]float64, nOutputSize)

    GenerateSeedNoise()
    GeneratePerlinNoise1D(  len(fPerlinNoise1D) , fNoiseSeed1D, fScalingBias,  fPerlinNoise1D  )

    nOutputWidth = screenW
    nOutputHeight = screenH
    fNoiseSeed2D = make([]float64, nOutputWidth*nOutputHeight )
    fPerlinNoise2D = make([]float64, nOutputWidth*nOutputHeight )

    GenerateSeedNoise2D()
    GeneratePerlinNoise2D( nOutputWidth, nOutputHeight , fNoiseSeed2D, fScalingBias ,  fPerlinNoise2D  )
}
func (self *MyView) Exit() {}
func (self *MyView) Update(t, dt float64) {

    graph.FillRect( self.screenImage, self.screenImage.Bounds() ,
                color.Black )

    if nMode == 1 {
        for x:=0; x<len(fPerlinNoise1D); x++ {
            // minus y effectively make y grows from bottom to top
            y := -int(fPerlinNoise1D[x] * float64(screenH/2)) + (screenH)/2
            graph.DrawLine( self.screenImage , x, y , x , screenH/2 , graph.COLOR_GREEN )
        }
    } else if nMode == 2 {
        for x:=0; x<nOutputWidth; x++ {
            for y:=0; y<nOutputHeight; y++ {
                gray := uint8(fPerlinNoise2D[ y*nOutputWidth + x ] * 255)
                self.screenImage.Set( x,y , color.Gray{ gray  } )
            }
        }
    } else {
        var col color.Color
        for x:=0; x<nOutputWidth; x++ {
            for y:=0; y<nOutputHeight; y++ {
                pixel_bw := uint8(fPerlinNoise2D[ y*nOutputWidth + x ] * 16)
                switch pixel_bw {
                case 0:
                    col = graph.COLOR_DARKBLUE
                case 1: fallthrough
                case 2: fallthrough
                case 3: fallthrough
                case 4:
                    col = color.RGBA{ 0,0,160 + pixel_bw * 16, 255 }

                case 5: fallthrough
                case 6: fallthrough
                case 7: fallthrough
                case 8:
                    col = color.RGBA{0,128 + (pixel_bw-4)*16  ,0,255} 

                case 9: fallthrough
                case 10: fallthrough
                case 11: fallthrough
                case 12:
                    g := 80 + (pixel_bw-8)*16
                    col = color.RGBA{g,g,g,255}
                case 13: fallthrough
                case 14: fallthrough
                case 15: fallthrough
                case 16:
                    g := 180 + (pixel_bw-8)*16
                    col = color.RGBA{g,g,g,255}

                } // end swith
                self.screenImage.Set(x,y, col )
            }
        } // end for

    }
}

func (self *MyView) OnKey(key glfw.Key) {
    switch key {
    case glfw.KeyR:
        if nMode == 1 {
            GenerateSeedNoise()
        } else {
            GenerateSeedNoise2D()
        }
    case glfw.KeyW:
        fScalingBias -= 0.1
        if fScalingBias <= 0.1 {
            fScalingBias = 0.1
        }
    case glfw.KeyS:
        fScalingBias += 0.1
    case glfw.Key1:
        nMode = 1
    case glfw.Key2:
        nMode = 2
    case glfw.Key3:
        nMode = 3
    default:
        return
    }
    if nMode == 1 {
        GeneratePerlinNoise1D(  len(fPerlinNoise1D) , fNoiseSeed1D, fScalingBias ,  fPerlinNoise1D  )
    } else {
        GeneratePerlinNoise2D( nOutputWidth, nOutputHeight , fNoiseSeed2D, fScalingBias ,  fPerlinNoise2D  )
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

// 1D
var fNoiseSeed1D []float64
var fPerlinNoise1D []float64

// 2D
var nOutputWidth , nOutputHeight int
var fNoiseSeed2D []float64
var fPerlinNoise2D []float64

var fScalingBias float64 = 0.5
var nMode int = 1

func GeneratePerlinNoise1D( nCount  int, fSeed []float64, fBias float64, fOutput []float64 ) {
    fOctaves := math.Log2(  float64(nCount ) )
    nOctaves := int(fOctaves)
    if fOctaves != float64(nOctaves) {
        panic( "output count must be power of 2" )
    }
    // fScaleAcc := 2 - 1/float64(nCount/2)
    for x:=0; x<nCount ; x++ {
        fNoise := 0.0
        fScale := 1.0
        fScaleAcc := 0.0
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
            fScaleAcc += fScale
            // halve scale for next octave
            fScale *= fBias
        }
        // output , make sure it lies between [0,1]
        fOutput[x] = fNoise / fScaleAcc
    }
}

func GeneratePerlinNoise2D( nWidth, nHeight  int, fSeed []float64, fBias float64, fOutput []float64 ) {
    fOctaves := math.Log2(  float64(nWidth ) )
    nOctaves := int(fOctaves)
    if fOctaves != float64(nOctaves) {
        panic( "output nWidth must be power of 2" )
    }
    // fScaleAcc := 2 - 1/float64(nCount/2)
    for x:=0; x<nWidth ; x++ {
        for y:=0; y<nWidth ; y++ {
            fNoise := 0.0
            fScale := 1.0
            fScaleAcc := 0.0
            for o:=0; o< nOctaves; o++ {
                nPitch := nWidth  >> o
                nSampleX1 := x &^(nPitch-1)  // X1,Y1; X2,Y2 is a rectangle
                nSampleY1 := y &^(nPitch-1)
                nSampleX2 := (nSampleX1 + nPitch) % nWidth
                nSampleY2 := (nSampleY1 + nPitch) % nWidth
                // how far into the pitch
                fBlendX   := float64(x - nSampleX1) / float64(nPitch)
                fBlendY   := float64(y - nSampleY1) / float64(nPitch)
                // interpolation    nSample1:0<---fBlend---> 1:nSample2
                // both these samples use the blendX  , because they are all the interpolation in x axis
                fSampleT := (1-fBlendX)*fSeed[nSampleY1*nWidth + nSampleX1] + fBlendX * fSeed[nSampleY1*nWidth + nSampleX2]
                fSampleB := (1-fBlendX)*fSeed[nSampleY2*nWidth + nSampleX1] + fBlendX * fSeed[nSampleY2*nWidth + nSampleX2]
                // interpolation in y-axis
                fSample := (1-fBlendY)*fSampleT + fBlendY*fSampleB

                // fNoise is going to accumulate the noise for location X within our output array. 
                fNoise += fSample * fScale
                fScaleAcc += fScale
                // halve scale for next octave
                fScale *= fBias
            }
            // output , make sure it lies between [0,1]
            fOutput[y*nWidth + x] = fNoise / fScaleAcc
        }
    }
}

func GenerateSeedNoise() {
    for i:=0; i<len( fNoiseSeed1D ); i++ {
        fNoiseSeed1D[i] = rand.Float64()
    }
}

func GenerateSeedNoise2D() {
    for i:=0; i< nOutputWidth*nOutputHeight; i++ {
        fNoiseSeed2D[i] = rand.Float64()
    }
}

