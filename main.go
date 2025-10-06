package main

import (
	"image"
	_ "image/png"
	"log"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var img *ebiten.Image
var imgs []*ebiten.Image
var difficulty int = 4
var width_seg int
var height_seg int

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("tux.png")
	if err != nil {
		log.Fatal(err)
	}

	width_seg = (img.Bounds().Max.X - img.Bounds().Min.X) / difficulty
	height_seg = (img.Bounds().Max.Y - img.Bounds().Min.Y) / difficulty

	imgs = make([]*ebiten.Image, (difficulty*difficulty) - 1)

	for i := 0; i < difficulty; i++ {
		for j := 0; j < difficulty; j++ {
			if i == difficulty - 1 && j == difficulty - 1 {
				break
			}
			rect := image.Rect(width_seg * j, height_seg * i, width_seg * (j + 1), height_seg * (i + 1))
			sub_image := img.SubImage(rect)
			imgs[i * difficulty + j] = ebiten.NewImageFromImage(sub_image)
		}
	}
}

type Game struct{}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		fmt.Println("Pressed A!")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		fmt.Println("Pressed D!")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		fmt.Println("Pressed W!")
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		fmt.Println("Pressed S!")
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

	for i := 0; i < difficulty; i++ {
		for j := 0; j < difficulty; j++ {
			if i == difficulty - 1 && j == difficulty - 1 {
				break
			}
			
			translate := ebiten.DrawImageOptions{}
			translate.GeoM.Translate(float64(width_seg * j + 10 * j), float64(height_seg * i + 10 * i))
			screen.DrawImage(imgs[i * difficulty + j], &translate)
		}
	}
}

func (g *Game) Layout(outsidej, outsidei int) (screenj, screeni int) {
	return img.Bounds().Max.X + 100, img.Bounds().Max.Y + 100
}

func main() {
	ebiten.SetWindowSize(img.Bounds().Max.X, img.Bounds().Max.Y)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}