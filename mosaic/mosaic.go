package mosaic

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

const mosaicRatio = 16

var (
	gophersImage *ebiten.Image
)

func init() {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(images.Gophers_jpg))
	if err != nil {
		log.Fatal(err)
	}
	gophersImage = ebiten.NewImageFromImage(img)
}

type Game struct {
	gophersRenderTarget *ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// shrink the image once
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1.0/mosaicRatio, 1.0/mosaicRatio)
	g.gophersRenderTarget.DrawImage(gophersImage, op)

	// Enlarge the shrunk image
	// filter is the nearest filter, so the result will be mosiac
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(mosaicRatio, mosaicRatio)
	screen.DrawImage(g.gophersRenderTarget, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func RunMosaic() {
	w, h := gophersImage.Bounds().Dx(), gophersImage.Bounds().Dy()
	g := &Game{
		gophersRenderTarget: ebiten.NewImage(w/mosaicRatio, h/mosaicRatio),
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Mosaic")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
