package sprites

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/ganim8/v2"
	"log"
	"time"
)

var (
	SkullAnimationSprite *ganim8.Sprite
	MenuFontSprite       *ganim8.Sprite
	BloodSprite          *ebiten.Image
)

func MustLoadSprites() {
	mustLoadSkullSheet()
	mustLoadMenuFontSheet()
	mustLoadBloodSprite()
}

func mustLoadBloodSprite() {
	spriteSheet, _, err := ebitenutil.NewImageFromFile("assets/sprites/blood.png")
	if err != nil {
		log.Panic(err)
	}

	BloodSprite = spriteSheet
}

func mustLoadSkullSheet() {
	//spriteSheet, _, err := ebitenutil.NewImageFromFile("assets/sprites/skull_spritesheet.png")
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//skullAnimationGrid := ganim8.NewGrid(16, 16, spriteSheet.Bounds().Dx(), spriteSheet.Bounds().Dy())
	//SkullAnimationSprite = ganim8.NewSprite(spriteSheet, skullAnimationGrid.Frames("1-2", "1-2"))
	spriteSheet, _, err := ebitenutil.NewImageFromFile("assets/sprites/skull_spritesheet2.png")
	if err != nil {
		log.Panic(err)
	}

	skullAnimationGrid := ganim8.NewGrid(32, 32, spriteSheet.Bounds().Dx(), spriteSheet.Bounds().Dy())
	SkullAnimationSprite = ganim8.NewSprite(spriteSheet, skullAnimationGrid.Frames("1-8", 1))

}

func mustLoadMenuFontSheet() {
	spriteSheet, _, err := ebitenutil.NewImageFromFile("assets/sprites/menufont_spritesheet.png")
	if err != nil {
		log.Panic(err)
	}

	spriteGrid := ganim8.NewGrid(15, 25, spriteSheet.Bounds().Dx(), spriteSheet.Bounds().Dy())
	MenuFontSprite = ganim8.NewSprite(spriteSheet, spriteGrid.Frames("1-12", "1-6"))
}

func NewSkullAnimation() *ganim8.Animation {
	return ganim8.NewAnimation(SkullAnimationSprite, time.Millisecond*120)
}
