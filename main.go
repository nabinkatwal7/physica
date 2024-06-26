package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	//"github.com/nabinkatwal7/physica/mosaic"
	"github.com/nabinkatwal7/physica/noise"
	//"github.com/nabinkatwal7/physica/hsv"
	//"github.com/nabinkatwal7/physica/infinitescroll"
	//github.com/nabinkatwal7/physica/filter"
	//"github.com/nabinkatwal7/physica/flood"
	//"github.com/nabinkatwal7/physica/font"
	//"github.com/nabinkatwal7/physica/blur"
	//"github.com/nabinkatwal7/physica/examples"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	//	ebiten.SetWindowSize(640, 480)
	//	ebiten.SetWindowTitle("Hello, World!")
	//	if err := ebiten.RunGame(&Game{}); err != nil {
	//		log.Fatal(err)
	//	}

	//examples.RunAnimation()
	//blur.RunBlur()
	//filter.RunFilter()
	//flood.RunFlood()
	//font.RunFont()
	//hsv.RunHSV()
	//infinitescroll.RunInfiniteScroll()
	//mosaic.RunMosaic()
	noise.RunNoise()
}
