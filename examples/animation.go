package examples

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

const (
	screenWidth  = 320
	screenHeight = 240

	frame0X     = 0
	frame0Y     = 32
	frameWidth  = 32
	frameHeight = 32
	frameCount  = 8
)

var (
	runnerImage *ebiten.Image
)

// implements ebiten.Game interface
type Game struct {
	count int
}

// Update proceeds game state. update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.count++
	return nil
}

// Draw draws the game screen. Draw is called every frame (typically 1/60[s] for 60Hz).
func (g *Game) Draw(screen *ebiten.Image) {
	// Drawing the animation
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	i := (g.count / 5) % frameCount
	sx, sy := frame0X+i*frameWidth, frame0Y
	screen.DrawImage(runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func RunAnimation() {
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}

	runnerImage = ebiten.NewImageFromImage(img)

	// specify the window size as you like. Here, a doubled size is specified
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Animation example")

	// call ebiten.RunGame to start your game loop
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
